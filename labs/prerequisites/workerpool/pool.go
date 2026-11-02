package workerpool

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Job is one independently handled unit of work.
type Job struct {
	ID string
}

// Result records one handler outcome. Err is a job-level error, not
// necessarily an error for the overall Run operation.
type Result struct {
	JobID string
	Value int
	Err   error
}

// Handler performs one job. It must cooperate with cancellation through ctx.
type Handler func(ctx context.Context, job Job) (int, error)

// Config bounds worker concurrency and queued admission. PerJobTimeout is
// applied to each handler invocation in a later exercise increment.
type Config struct {
	Workers       int
	QueueCapacity int
	PerJobTimeout time.Duration
}

func validateConfig(cfg Config) error {
	if cfg.Workers <= 0 {
		return errors.New("workers must be > 0")
	}
	if cfg.QueueCapacity < 0 {
		return errors.New("queueCapacity must be >= 0")
	}
	if cfg.PerJobTimeout <= 0 {
		return errors.New("perJobTimeout must be > 0")
	}
	return nil
}

// Run executes the bounded, cancellable worker-pool contract in README.md.
func Run(
	ctx context.Context,
	cfg Config,
	jobs []Job,
	handle Handler,
) ([]Result, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	jobQueue := make(chan Job, cfg.QueueCapacity)
	results := make(chan Result, len(jobs))
	var wg sync.WaitGroup

	for i := 0; i < cfg.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobQueue:
					if !ok {
						return
					}
					if ctx.Err() != nil {
						return
					}
					jobCtx, cancel := context.WithTimeout(ctx, cfg.PerJobTimeout)
					value, err := handle(jobCtx, job)
					cancel()
					select {
					case <-ctx.Done():
						return
					case results <- Result{
						JobID: job.ID,
						Value: value,
						Err:   err,
					}:

					}
				}
			}
		}()
	}

	go func() {
		for _, job := range jobs {
			select {
			case <-ctx.Done():
				close(jobQueue)
				return
			case jobQueue <- job:
			}
		}
		close(jobQueue)

	}()

	wg.Wait()
	close(results)

	var finalResults []Result
	for r := range results {
		finalResults = append(finalResults, r)
	}
	return finalResults, ctx.Err()
}
