# Worker Pool Exercise and Review Process

## Objective

Implement a bounded, cancellable Go worker pool that satisfies the contract in
[README.md](README.md). This is a prerequisite lab, not Phase 1 product code.

## Task cards

Each increment is assigned with its path, concept, learner-owned code boundary,
acceptance criteria, verification command, and explicit exclusions. Do not
begin a later card until the previous implementation and the agent-authored
review tests have been discussed.

### Card 1 — ordinary job-error isolation

- **Path:** `labs/prerequisites/workerpool/pool.go`
- **Concept:** distinguish a per-job handler failure from an operation-wide
  failure.
- **Why it matters:** a dispatcher/controller must continue independent work
  after one item fails, while preserving the failed item's diagnostic.
- **Learner-owned scope:** replace the `Run` TODO with sequential processing
  that calls `handle(ctx, job)` for every job and returns one `Result` per job.
- **Acceptance criteria:** input order is preserved; each result preserves job
  ID, value, and handler error; one ordinary handler error does not stop later
  jobs; `Run` itself returns a nil operation error for that case.
- **Verification:** after the learner submits the code, the agent adds and runs
  deterministic unit tests with `go test ./labs/prerequisites/workerpool`.
  A green result proves ordinary error isolation only.
- **Out of scope:** goroutines, bounded queues, worker count, cancellation,
  deadlines, race safety, retries, persistence, and throughput.

### Card 2 — bounded concurrency and orderly completion

- **Path:** `labs/prerequisites/workerpool/pool.go`
- **Scaffold:** the Card 2 TODO block directly inside `Run` in `pool.go`; fill
  it in by evolving the verified Card 1 baseline in place.
- **Concept:** bound active handler calls and queued work while ensuring every
  accepted job reaches one result before `Run` returns.
- **Why it matters:** control-plane workers must cap CPU, memory, and upstream
  pressure. Unbounded goroutines or queues turn overload into a resource
  outage; premature channel closure turns it into a panic or lost result.
- **Learner-owned scope:** replace Card 1's sequential loop with `cfg.Workers`
  worker goroutines and an internal jobs channel buffered by
  `cfg.QueueCapacity`. Make `Run` the sole sender/closer of jobs; wait for all
  workers before closing any shared result path and returning the collected
  results. Validate positive workers, non-negative queue capacity, and a
  positive per-job timeout. Continue to pass the parent context directly to
  handlers and keep ordinary handler errors in `Result.Err`.
- **Acceptance criteria:** no more than `cfg.Workers` handlers are active at
  once; every input job has exactly one result; no worker closes a shared
  channel; invalid configuration returns an operation error without invoking
  the handler; an ordinary job error still does not halt independent work.
- **Verification:** after the learner submits the code, the agent adds tests
  with channel-based coordination (not sleeps) for the active-handler bound,
  all-job completion, and invalid configuration; then runs
  `go test ./labs/prerequisites/workerpool`. A green result proves the stated
  in-process bound, not cancellation or race freedom.
- **Ordering note:** with one worker, Card 1's `a`, `b`, `c` execution order
  remains observable. With multiple workers, completion/result order is not an
  API guarantee unless a later requirement explicitly introduces indexed
  ordered collection.
- **Out of scope:** selecting on `ctx.Done()`, abandoning queued jobs on
  cancellation, per-job child deadline contexts, retry policy, panic recovery,
  persistence, cross-process distribution, and race hardening.

### Card 3 — cooperative cancellation and shutdown

- **Path:** `labs/prerequisites/workerpool/pool.go`
- **Scaffold:** the Card 3 TODO block directly inside `Run`; evolve the
  demonstrated Card 2 baseline in place.
- **Concept:** cancellation is a cooperative signal at every potentially
  blocking producer/worker/result boundary; it does not forcibly stop a
  handler goroutine.
- **Why it matters:** shutdown, leadership loss, request deadlines, and retry
  abandonment must release worker resources without presenting unstarted work
  as successful or stranding goroutines behind a blocked channel operation.
- **Learner-owned scope:** add `ctx.Done()` alternatives to job admission,
  worker job receive, and result send. On parent cancellation, stop admission
  and abandon queued/unstarted jobs; pass the parent context to started
  handlers; return partial results plus `ctx.Err()`. Keep producer-owned jobs
  closure, coordinator-owned results closure, and single-owner result
  collection. Do not derive child timeout contexts yet.
- **Acceptance criteria:** a cancel after one handler starts completes without
  a wall-clock sleep; queued work does not start; `Run` returns
  `context.Canceled` (or the parent context's terminal error) plus partial
  results; all workers terminate; ordinary job errors remain isolated.
- **Verification:** after the learner submits the code, the agent adds a
  channel-coordinated cancellation test and checks it with
  `go test ./labs/prerequisites/workerpool`. Race/repetition hardening remains
  a later card.
- **Out of scope:** per-job `WithTimeout`, retry/backoff, panic recovery,
  persistent queues, exactly-once claims, external I/O, and performance work.

### Card 4 — per-job deadline propagation

- **Path:** `labs/prerequisites/workerpool/pool.go`
- **Scaffold:** the Card 4 TODO block directly inside `Run`; evolve the
  demonstrated Card 3 baseline in place.
- **Primary resource:** Go [`context.WithTimeout`](https://pkg.go.dev/context#WithTimeout)
  documentation — **required before practice**.
- **Concept:** a parent operation context controls pool cancellation; a child
  context controls one handler's time budget. These are different error scopes.
- **Why it matters:** control-plane workers need a bounded attempt duration so
  one slow reconciliation or upstream request cannot consume capacity forever.
- **Learner-owned scope:** derive `jobCtx` with `context.WithTimeout(ctx,
  cfg.PerJobTimeout)` immediately before `handle`; call the returned cancel
  function immediately after it returns, not with a long-lived `defer`. Retain
  every Card 3 channel/cancellation rule. A handler that returns because its
  child deadline expires contributes `context.DeadlineExceeded` in that job's
  result while `Run` remains nil when the parent context is still active.
- **Acceptance criteria:** a handler that waits on its context gets a deadline;
  its result has `context.DeadlineExceeded`; `Run` returns a nil operation
  error; the test uses no sleep for synchronization; no cancellation behavior
  regresses.
- **Verification:** after the learner submits the code, the agent adds a
  deadline-propagation test and runs `go test ./labs/prerequisites/workerpool`.
- **Out of scope:** retry/backoff, deadline selection policy beyond the one
  config value, panic recovery, persistence, distributed leases, and profiling.

### Card 5 — hardening and limitation record

- **Path:** `labs/prerequisites/workerpool/` (implementation and generated
  tests)
- **Concept:** repeated race-enabled verification increases confidence that
  current deterministic tests do not conceal a schedule-dependent defect; it
  does not prove all possible schedules or arbitrary handler behavior.
- **Why it matters:** control-plane concurrency bugs often occur only under
  particular interleavings. A one-time green run is weaker evidence than a
  repeated race-enabled run with stated limitations.
- **Learner-owned scope:** review the current `Run` lifecycle and explain the
  limitation that a handler ignoring its context can delay clean shutdown. Do
  not add retries, panic recovery, persistence, or production metrics here.
- **Acceptance criteria:** `go test -race ./labs/prerequisites/workerpool
  -count=50 -timeout=10s` passes; the learner distinguishes that evidence from
  a proof of absence of races; the handler-cooperation limitation is recorded.
- **Verification:** agent runs the exact command and records the result. No
  external resource is needed for this small verification task.
- **Out of scope:** performance benchmarking, fuzzing, model checking,
  distributed execution, and a claim of universal race freedom.

## Before coding

For Card 2, read the Go [`sync.WaitGroup` documentation](https://pkg.go.dev/sync#WaitGroup)
and the Tour of Go's [Channels](https://go.dev/tour/concurrency/2) and
[Buffered Channels](https://go.dev/tour/concurrency/3) sections. They are
required before practice. `pool.go` contains the verified Card 1 baseline and
the Card 2 lifecycle slots. Fill the TODOs by evolving `Run` in place; do not
add cancellation or deadline behavior. Then submit the implementation diff.
The agent reviews the code and adds/shows the unit-test diff, including edge
cases, before the next increment begins.

## Incremental exercise sequence

1. **Configuration and ordinary errors**
   - Add Config, Job, Result, Handler, and Run.
   - Implement one job error as a result, with later independent work still
     running; the agent adds review tests.
   - Add invalid-config and all-success coverage during the following review.
2. **Bounded concurrency**
   - Add the bounded jobs channel and worker count.
   - Test that no more than Workers handlers run concurrently using a
     mutex-protected counter and a release channel, not wall-clock sleeps.
3. **Cancellation and shutdown**
   - Add context-aware producer send, worker receive, and result send.
   - Add the cancellation-abandons-queued-jobs test using started/release
     channels for deterministic coordination.
   - Ensure the producer closes jobs, and one coordinator closes results after
     the worker wait group completes.
4. **Per-job deadline**
   - Derive the child context in the worker immediately before the handler.
   - Call the child cancel function after the handler returns.
   - Add the deadline-propagation test.
5. **Hardening**
   - Run race detection and repeated tests.
   - Record any remaining limitation: a handler that ignores its context can
     delay clean shutdown.

## Learner self-review checklist

Before requesting review, confirm each item against the diff:

- [ ] Every channel has one closing owner; no worker closes the shared results
  channel.
- [ ] Every blocking job send, job receive, and result send has a cancellation
  alternative.
- [ ] The producer stops admission after cancellation and closes jobs exactly
  once.
- [ ] The result collector terminates after workers finish; no goroutine waits
  on an unread public result channel.
- [ ] Child timeout contexts are derived per job and canceled after handler
  return, not deferred across the worker loop.
- [ ] Ordinary handler errors do not cancel independent jobs.
- [ ] Cancellation returns partial results plus the operation context error.
- [ ] Tests coordinate with channels, wait groups, or contexts, not arbitrary
  sleeps.

## Review gate

For each increment, provide:

1. The learner implementation diff and an explanation of its invariant.
2. The agent-generated test diff, its invariant, and the failure it is meant
   to detect.
3. Exact command output for go test ./labs/prerequisites/workerpool.
4. At hardening, exact command output for go test -race ./... and
   go test -race ./labs/prerequisites/workerpool -count=50 -timeout=10s.

The reviewer checks the contract, channel ownership, safety and liveness
separately, deterministic tests, race output, and unverified assumptions. A
passing test is evidence only for its stated property; it does not prove
arbitrary handler cancellation or crash recovery.

## Completion evidence

The lab is ready for reassessment only after all acceptance tests and both
hardening commands pass, the learner can explain the cancellation/result
contract without notes, and the limitations are recorded in progress files.
