# Current Roadmap State

> This is the human-readable resume checkpoint. Keep it factual and commit/push updates with the work they describe. Git history and upstream status remain authoritative for whether a checkpoint is shared.

## Status
- **Checkpoint status:** IN PROGRESS — targeted prerequisite block underway
- **Active phase:** Prerequisite block
- **Active week/task:** Worker-pool prerequisite Exercise 1 — write and review the ordinary-job-error test before implementation; do not begin Phase 1 yet
- **Last verified checkpoint:** 2026-11-02 career-aligned roadmap and design-only worker-pool lab checkpoint, committed and pushed; no behavior implementation exists
- **Last verified commit:** 8eb9ed6 on baseline_tested
- **Remote sync:** baseline_tested matched origin/baseline_tested after pushing 8eb9ed6 on 2026-11-02.

## Last completed study checkpoint
- **Date:** 2026-10-20
- **Objective:** Complete focused Linux, networking, and OS-memory prerequisite evidence; keep revision material mapped to the authoritative roadmap.
- **Completed work:** Completed a disposable-container Linux HTTP-service triage: listener/PID, FD ownership, FD table, logs, request syscalls, idle readiness waits, and Docker port mapping. Captured and interpreted a bounded HTTPS `curl -v` request through DNS, TCP, TLS, HTTP/2, and `200`. Completed page-cache/RSS/reclaim reasoning plus a bounded 64 MiB Docker file-read observation and cleanup. Added revision-note/cheatsheet material and an automated Phase 1–7 title/order alignment check for the roadmap map and core-phase notes.
- **Evidence:** Linux is **Practiced**: the traced request accepted on listener FD 3, read 78 request bytes, logged 61 bytes, sent 156-byte headers and an 805-byte body, and shut down without displayed errno; idle `ppoll` behavior was observed. Networking is **Demonstrated for baseline scope**: a verbose HTTPS request resolved two IPv4 addresses, connected to port 443, verified TLS 1.3 certificate details, negotiated HTTP/2, sent `GET /`, and received `HTTP/2 200`. OS/memory is **Practiced**: RSS/cache/reclaim model was explained; the 64 MiB observation showed cache counters but both reads rounded to `0.00s`, so no speed claim was made. Kernel OOM logs are **Not assessed** in Docker Desktop because Docker Debug received `Operation not permitted` from `dmesg`.
- **Tests/commands actually run:** See the 2026-10-20 session entry in `progress/weekly-log.md` for the exact learner commands and outcomes. Agent verification: `./scripts/check-revision-roadmap-alignment.sh` passed; `git diff --check` passed. Go tests/race tests: **not run**. Benchmarks: **not run**. Graphify refresh: **skipped/not run** because no semantic backend was explicitly approved and available; no graph artifacts were changed or staged.
- **Open gaps/blockers:** Targeted prerequisite block remains required before Phase 1. Linux result interpretation needs one later fresh scenario; OS kernel OOM log evidence needs a dedicated Linux VM; Go worker-pool runtime evidence, Kubernetes finalizer semantics, cloud IAM/cost, Git bisection, and container CI smoke testing remain open. The revision-alignment check only guards Phase 1–7 title/order and map coverage; the roadmap remains the authority for all curriculum scope.

## Resume reconciliation

- **Date:** 2026-11-02
- **Repository evidence:** The active branch is baseline_tested; git fetch --prune succeeded and HEAD 414a1a2 matched its upstream. The prior state file's reference to f35419b was stale.
- **Local work preserved:** Modified Go-assessment/revision files plus untracked go.mod and labs/prerequisites/workerpool form the intended next prerequisite exercise. The lab has a design specification and review process; no worker-pool implementation or behavior test exists.
- **Decision:** Learner confirmed that this local working tree was the authoritative continuation. It was captured in commit 8eb9ed6 and pushed to origin/baseline_tested on 2026-11-02.

## Resume here next session
1. Inspect Git status, latest commit and upstream sync using the start-of-session protocol in `AGENTS.md`.
2. Complete Exercise 1: write TestRunContinuesAfterJobError in the worker-pool lab, review its property, then implement only enough behavior to pass it.
3. Continue the later cancellation/deadline/race increments only after that review.
4. Update progress records only from observed evidence.

## Next action
Write and review TestRunContinuesAfterJobError before implementing the first worker-pool increment.

## Project-studio scaffold

- **Date:** 2026-11-02
- **Work completed:** Added a roadmap-mapped, test-first exercise catalog for
  process monitoring, dispatcher, rate limiting, cache behavior, mini event
  log, outbox/CDC, scheduler, gateway policy, log aggregation, and a
  role-gated storage lab. The Infrastructure Lifecycle Control Plane remains
  the only progressive capstone.
- **Evidence status:** Documentation/scaffolding only. No studio exercise has
  implementation, test, benchmark, or learner demonstration evidence.
- **Guard:** Do not start a studio item outside its mapped phase. The active
  prerequisite remains the worker-pool first-test review.

## Career-alignment update

- **Date:** 2026-11-02
- **Work completed:** Merged the supplied career-alignment refinement into the
  authoritative roadmap without replacing the existing phase sequence or
  prerequisite evidence. Added a career scorecard, application tracker,
  validation prompt, ADR, and mapped revision notes.
- **Evidence:** Phase-title/order guard and whitespace check passed after the
  merge. The root Go module check passed for the design-only worker-pool
  package with no test files. No code behavior, benchmark, cloud action, or
  prerequisite exit test was run by this roadmap update. On 2026-11-02,
  `ROADMAP_2030_REVISED_REFERENCE.md` in the supplied kit was verified as
  identical to the repository roadmap immediately before the career-alignment
  merge; the ADR now records it as the technical-curriculum baseline.
- **Active effect:** Primary direction is storage/workload control planes;
  AI specialization is conditional on evidence. The active task remains the
  first Go worker-pool test.

## Shared checkpoint

- **Commit:** 8eb9ed6 — docs(roadmap): align career direction and add worker-pool lab
- **Upstream:** Pushed to origin/baseline_tested on 2026-11-02.
- **Graph artifact:** Skipped by the pre-commit hook because no supported
  Graphify semantic backend is configured; no graph artifacts were refreshed.

## Decision log
- Keep `ROADMAP_2030_REVISED.md` as the curriculum source of truth.
- Keep Go/control-plane/distributed systems as the core path; Rust remains a sequenced follow-up unless explicitly requested earlier.
- Keep phase revision notes and cheatsheets under `revision/`; `revision/ROADMAP_MAP.md` maps them to roadmap sections and defines the agent update contract.
- Enforce that `revision/ROADMAP_MAP.md` and `revision/notes/core-phases.md` retain the roadmap's Phase 1–7 titles/order with `scripts/check-revision-roadmap-alignment.sh`, which the pre-commit hook runs before the optional Graphify step.
