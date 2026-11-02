# Bounded Cancellable Worker Pool Lab

**Roadmap:** [Phase 1, Week 2](../../../ROADMAP_2030_REVISED.md#weekly-sequence)  
**Status:** Design only — no worker-pool behavior is implemented or tested yet.

This prerequisite lab isolates Go cancellation, channel ownership, and shutdown
semantics. It is not the Phase 1 reliable job dispatcher. The dispatcher starts
only after prerequisite placement allows Phase 1 to begin.

## Requirements

- The pool has a positive worker count and a bounded internal job queue.
- The producer is the sole owner of the jobs-channel close operation.
- A coordinator closes results only after every worker has exited.
- Every potentially blocking job receive, job send, and result send selects on
  context cancellation.
- Each handler receives a per-job child context derived with WithTimeout.
  The worker calls that child context's cancel function when the handler
  returns; the handler is responsible for cooperative cancellation.
- Parent cancellation stops admission and abandons queued, unstarted jobs.
- A normal job error is returned in that job's result and does not stop other
  independent jobs.
- On parent cancellation, Run returns partial results plus the context error.
  Missing results mean incomplete or abandoned work, never success.

## Target API

The first implementation uses this deliberately small API:

    type Job struct {
        ID string
    }

    type Result struct {
        JobID string
        Value int
        Err   error
    }

    type Handler func(context.Context, Job) (int, error)

    type Config struct {
        Workers       int
        QueueCapacity int
        PerJobTimeout time.Duration
    }

    func Run(
        ctx context.Context,
        cfg Config,
        jobs []Job,
        handle Handler,
    ) ([]Result, error)

Run owns internal channels and collects results before returning. This lab does
not expose a public results channel. That keeps the cancellation contract
testable while retaining internal result-send cancellation behavior.

## Assumptions and failure model

- handle is expected to return promptly after its context is canceled. A
  handler that ignores its context can delay pool shutdown; Go cannot forcibly
  stop that goroutine safely.
- The lab does not persist jobs, promise exactly-once execution, recover from a
  process crash, or expose a network API.
- Result collection uses memory proportional to completed results. Bounded job
  admission prevents unbounded queued work, not unbounded caller input slices.

## Acceptance tests to write before implementation

1. TestRunContinuesAfterJobError: use Workers 1, jobs a, b, and c, and a
   handler that returns errors.New("boom") only for b. Assert that Run returns
   a nil operation error, produces three results, reports boom only for b, and
   invokes the handler for c after b. This is the first test to write.
2. TestRunRejectsInvalidConfig: reject zero workers, negative queue capacity,
   and a non-positive per-job timeout.
3. TestRunReturnsAllSuccessfulResults: every submitted job gets exactly one
   successful result; no duplicate job IDs appear.
4. TestRunCancellationAbandonsQueuedJobs: one worker starts a handler that
   waits on its context. Cancel only after that handler signals it started.
   Assert that Run returns a context-canceled operation error, no queued job
   begins, and the test completes without waiting on a fixed sleep.
5. TestRunPropagatesPerJobDeadline: the handler waits for its job context to
   finish. Assert its result has context.DeadlineExceeded while the overall
   operation is not canceled.
6. go test -race ./... passes after behavior tests exist.

## Verification status

Module/package scaffolding only. On 2026-10-21, go test ./... passed for this
package with no test files. It cannot verify any requirement until tests and
behavior are implemented. Race tests are not run.

See [REVIEW.md](REVIEW.md) for the mandatory learner and reviewer gates.
