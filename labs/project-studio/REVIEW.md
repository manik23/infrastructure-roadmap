# Project Studio Review Gate

Apply this gate before each increment. A green test only establishes the
property that test actually checks.

## Learner submission

Provide:

1. The exercise number and increment being attempted.
2. Requirements, assumptions, resource bound, security boundary, and
   safety/liveness invariant as applicable.
3. The first test, model, or design contract before implementation.
4. The changed files and an explanation of the learner-owned logic.
5. Exact commands and output for tests, race tests, benchmarks, or fault
   checks that were actually run; label all others `not run`.
6. One known limitation or failure case.

## Reviewer checks

- The exercise is at or before the active roadmap phase and its prerequisite
  gate is satisfied.
- Channel/lock/state ownership and cancellation are explicit where relevant.
- Time, retries, duplicates, partial failure, resource bounds, and cleanup
  are tested or explicitly deferred.
- A miniature implementation does not make unsupported production claims.
- Tests use deterministic coordination instead of arbitrary sleeps when
  concurrency is involved.
- Progress and revision notes distinguish scaffolded, implemented, tested,
  and demonstrated work.
