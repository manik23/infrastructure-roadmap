package workerpool

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunContinuesAfterJobError(t *testing.T) {
	t.Parallel()

	boom := errors.New("boom")
	jobs := []Job{{ID: "a"}, {ID: "b"}, {ID: "c"}}
	var calls []string

	results, err := Run(
		context.Background(),
		validTestConfig(),
		jobs,
		func(_ context.Context, job Job) (int, error) {
			calls = append(calls, job.ID)
			if job.ID == "b" {
				return 20, boom
			}
			return len(calls) * 10, nil
		},
	)

	if err != nil {
		t.Fatalf("Run() error = %v, want nil for an ordinary job error", err)
	}
	if want := []string{"a", "b", "c"}; !reflect.DeepEqual(calls, want) {
		t.Fatalf("handler call order = %v, want %v", calls, want)
	}
	if len(results) != len(jobs) {
		t.Fatalf("result count = %d, want %d", len(results), len(jobs))
	}

	for index, wantID := range []string{"a", "b", "c"} {
		if results[index].JobID != wantID {
			t.Errorf("results[%d].JobID = %q, want %q", index, results[index].JobID, wantID)
		}
	}
	if !errors.Is(results[1].Err, boom) {
		t.Errorf("results[1].Err = %v, want boom", results[1].Err)
	}
	if results[0].Err != nil || results[2].Err != nil {
		t.Errorf("successful job errors = (%v, %v), want (nil, nil)", results[0].Err, results[2].Err)
	}
	if results[0].Value != 10 || results[1].Value != 20 || results[2].Value != 30 {
		t.Errorf("result values = (%d, %d, %d), want (10, 20, 30)", results[0].Value, results[1].Value, results[2].Value)
	}
}

func TestRunPropagatesCallerContextValuesToEveryHandler(t *testing.T) {
	t.Parallel()

	type contextKey struct{}
	key := contextKey{}
	ctx := context.WithValue(context.Background(), key, "caller-context")
	jobs := []Job{{ID: "a"}, {ID: "b"}}
	var values []any

	_, err := Run(ctx, validTestConfig(), jobs, func(got context.Context, _ Job) (int, error) {
		values = append(values, got.Value(key))
		return 0, nil
	})
	if err != nil {
		t.Fatalf("Run() error = %v, want nil", err)
	}
	if len(values) != len(jobs) {
		t.Fatalf("handler calls = %d, want %d", len(values), len(jobs))
	}
	for index, got := range values {
		if got != "caller-context" {
			t.Errorf("handler context value %d = %v, want caller-context", index, got)
		}
	}
}

func TestRunWithNoJobsReturnsNoResults(t *testing.T) {
	t.Parallel()

	called := false
	results, err := Run(context.Background(), validTestConfig(), nil, func(context.Context, Job) (int, error) {
		called = true
		return 0, nil
	})
	if err != nil {
		t.Fatalf("Run() error = %v, want nil", err)
	}
	if called {
		t.Fatal("handler was called for no jobs")
	}
	if len(results) != 0 {
		t.Fatalf("result count = %d, want 0", len(results))
	}
}

func TestRunBoundsConcurrentHandlersAndReturnsEveryJob(t *testing.T) {
	t.Parallel()

	const workerCount = 2
	jobs := []Job{{ID: "a"}, {ID: "b"}, {ID: "c"}, {ID: "d"}}
	started := make(chan struct{}, workerCount)
	release := make(chan struct{})
	var releaseOnce sync.Once
	t.Cleanup(func() { releaseOnce.Do(func() { close(release) }) })

	var active, maxActive int
	var activeMu sync.Mutex
	completed := make(chan struct{})
	var results []Result
	var runErr error

	go func() {
		results, runErr = Run(
			context.Background(),
			Config{Workers: workerCount, QueueCapacity: 1, PerJobTimeout: time.Second},
			jobs,
			func(context.Context, Job) (int, error) {
				activeMu.Lock()
				active++
				if active > maxActive {
					maxActive = active
				}
				activeMu.Unlock()

				started <- struct{}{}
				<-release

				activeMu.Lock()
				active--
				activeMu.Unlock()
				return 1, nil
			},
		)
		close(completed)
	}()

	for range workerCount {
		select {
		case <-started:
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for all workers to start")
		}
	}

	activeMu.Lock()
	gotMaxActive := maxActive
	activeMu.Unlock()
	if gotMaxActive != workerCount {
		t.Fatalf("maximum active handlers = %d, want %d", gotMaxActive, workerCount)
	}

	releaseOnce.Do(func() { close(release) })
	select {
	case <-completed:
	case <-time.After(time.Second):
		t.Fatal("Run did not return after handlers were released")
	}

	if runErr != nil {
		t.Fatalf("Run() error = %v, want nil", runErr)
	}
	if len(results) != len(jobs) {
		t.Fatalf("result count = %d, want %d", len(results), len(jobs))
	}
	seen := make(map[string]int, len(results))
	for _, result := range results {
		seen[result.JobID]++
		if result.Err != nil {
			t.Errorf("result for %q has error = %v, want nil", result.JobID, result.Err)
		}
	}
	for _, job := range jobs {
		if seen[job.ID] != 1 {
			t.Errorf("result count for job %q = %d, want 1", job.ID, seen[job.ID])
		}
	}
}

func TestRunRejectsInvalidConfigWithoutCallingHandler(t *testing.T) {
	t.Parallel()

	configs := map[string]Config{
		"zero workers":            {Workers: 0, QueueCapacity: 1, PerJobTimeout: time.Second},
		"negative queue capacity": {Workers: 1, QueueCapacity: -1, PerJobTimeout: time.Second},
		"zero per-job timeout":    {Workers: 1, QueueCapacity: 1, PerJobTimeout: 0},
	}
	for name, cfg := range configs {
		t.Run(name, func(t *testing.T) {
			called := false
			_, err := Run(context.Background(), cfg, []Job{{ID: "a"}}, func(context.Context, Job) (int, error) {
				called = true
				return 0, nil
			})
			if err == nil {
				t.Fatal("Run() error = nil, want invalid configuration error")
			}
			if called {
				t.Fatal("handler was called for invalid configuration")
			}
		})
	}
}

func TestRunCancellationAbandonsQueuedJobs(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	started := make(chan struct{})
	completed := make(chan struct{})
	var results []Result
	var runErr error
	var queuedJobStarted atomic.Bool

	go func() {
		results, runErr = Run(
			ctx,
			Config{Workers: 1, QueueCapacity: 1, PerJobTimeout: time.Second},
			[]Job{{ID: "running"}, {ID: "queued-1"}, {ID: "queued-2"}},
			func(ctx context.Context, job Job) (int, error) {
				if job.ID != "running" {
					queuedJobStarted.Store(true)
					return 0, nil
				}
				close(started)
				<-ctx.Done()
				return 0, ctx.Err()
			},
		)
		close(completed)
	}()

	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for the first handler to start")
	}
	cancel()

	select {
	case <-completed:
	case <-time.After(time.Second):
		t.Fatal("Run did not return after cancellation")
	}
	if !errors.Is(runErr, context.Canceled) {
		t.Fatalf("Run() error = %v, want context.Canceled", runErr)
	}
	if queuedJobStarted.Load() {
		t.Fatal("a queued job began after cancellation")
	}
	if len(results) > 1 {
		t.Fatalf("result count = %d, want at most one partial running-job result", len(results))
	}
}

func TestRunAlreadyCanceledDoesNotStartHandlers(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	completed := make(chan struct{})
	var results []Result
	var runErr error
	var handlerCalled atomic.Bool

	go func() {
		results, runErr = Run(
			ctx,
			Config{Workers: 2, QueueCapacity: 0, PerJobTimeout: time.Second},
			[]Job{{ID: "a"}, {ID: "b"}},
			func(context.Context, Job) (int, error) {
				handlerCalled.Store(true)
				return 0, nil
			},
		)
		close(completed)
	}()

	select {
	case <-completed:
	case <-time.After(time.Second):
		t.Fatal("Run did not return for an already-canceled context")
	}
	if !errors.Is(runErr, context.Canceled) {
		t.Fatalf("Run() error = %v, want context.Canceled", runErr)
	}
	if handlerCalled.Load() {
		t.Fatal("handler was called for an already-canceled context")
	}
	if len(results) != 0 {
		t.Fatalf("result count = %d, want 0", len(results))
	}
}

func TestRunPropagatesPerJobDeadline(t *testing.T) {
	t.Parallel()

	type outcome struct {
		results []Result
		err     error
	}
	completed := make(chan outcome, 1)
	var handlerSawDeadline atomic.Bool
	go func() {
		results, err := Run(
			context.Background(),
			Config{Workers: 1, QueueCapacity: 1, PerJobTimeout: 20 * time.Millisecond},
			[]Job{{ID: "deadline"}},
			func(jobCtx context.Context, _ Job) (int, error) {
				_, hasDeadline := jobCtx.Deadline()
				handlerSawDeadline.Store(hasDeadline)
				<-jobCtx.Done()
				return 0, jobCtx.Err()
			},
		)
		completed <- outcome{results: results, err: err}
	}()

	select {
	case got := <-completed:
		if got.err != nil {
			t.Fatalf("Run() error = %v, want nil for a child deadline", got.err)
		}
		if len(got.results) != 1 {
			t.Fatalf("result count = %d, want 1", len(got.results))
		}
		if !errors.Is(got.results[0].Err, context.DeadlineExceeded) {
			t.Fatalf("result error = %v, want context.DeadlineExceeded", got.results[0].Err)
		}
	case <-time.After(time.Second):
		t.Fatal("Run did not return after the per-job deadline")
	}
	if !handlerSawDeadline.Load() {
		t.Fatal("handler context has no deadline")
	}
}

func validTestConfig() Config {
	return Config{Workers: 1, QueueCapacity: 1, PerJobTimeout: time.Second}
}
