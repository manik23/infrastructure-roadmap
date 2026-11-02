# Worker Pool Exercise and Review Process

## Objective

Implement a bounded, cancellable Go worker pool that satisfies the contract in
[README.md](README.md). This is a prerequisite lab, not Phase 1 product code.

## Before coding

Required reading: the Go [context package overview](https://pkg.go.dev/context)
and the documentation for WithCancel and WithTimeout.

Write pool_test.go first. Start only with TestRunContinuesAfterJobError; do not
implement cancellation behavior until that test is reviewed and passes.

## Incremental exercise sequence

1. **Configuration and ordinary errors**
   - Add Config, Job, Result, Handler, and Run.
   - Make the first test pass: one job error is a result, and job c still runs
     after job b fails.
   - Add invalid-config and all-success tests.
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

1. The changed files or a diff.
2. The test name and what property it proves.
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
