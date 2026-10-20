# Current Roadmap State

> This is the human-readable resume checkpoint. Keep it factual and commit/push updates with the work they describe. Git history and upstream status remain authoritative for whether a checkpoint is shared.

## Status
- **Checkpoint status:** IN PROGRESS — targeted prerequisite block underway
- **Active phase:** Prerequisite block
- **Active week/task:** Define bounded Go worker-pool cancellation, closure, and result-ownership semantics; do not begin Phase 1 yet
- **Last verified checkpoint:** 2026-10-20 prerequisite evidence captured for Linux triage, happy-path HTTPS, and page-cache/reclaim reasoning; records are local and uncommitted
- **Last verified commit:** `f35419b` on `baseline_tested` (previously documented as matching `origin/baseline_tested`)
- **Remote sync:** Not re-verified this session. The current local records are not committed or pushed.

## Last session summary
- **Date:** 2026-10-20
- **Objective:** Complete focused Linux, networking, and OS-memory prerequisite evidence; keep revision material mapped to the authoritative roadmap.
- **Completed work:** Completed a disposable-container Linux HTTP-service triage: listener/PID, FD ownership, FD table, logs, request syscalls, idle readiness waits, and Docker port mapping. Captured and interpreted a bounded HTTPS `curl -v` request through DNS, TCP, TLS, HTTP/2, and `200`. Completed page-cache/RSS/reclaim reasoning plus a bounded 64 MiB Docker file-read observation and cleanup. Added revision-note/cheatsheet material and an automated Phase 1–7 title/order alignment check for the roadmap map and core-phase notes.
- **Evidence:** Linux is **Practiced**: the traced request accepted on listener FD 3, read 78 request bytes, logged 61 bytes, sent 156-byte headers and an 805-byte body, and shut down without displayed errno; idle `ppoll` behavior was observed. Networking is **Demonstrated for baseline scope**: a verbose HTTPS request resolved two IPv4 addresses, connected to port 443, verified TLS 1.3 certificate details, negotiated HTTP/2, sent `GET /`, and received `HTTP/2 200`. OS/memory is **Practiced**: RSS/cache/reclaim model was explained; the 64 MiB observation showed cache counters but both reads rounded to `0.00s`, so no speed claim was made. Kernel OOM logs are **Not assessed** in Docker Desktop because Docker Debug received `Operation not permitted` from `dmesg`.
- **Tests/commands actually run:** See the 2026-10-20 session entry in `progress/weekly-log.md` for the exact learner commands and outcomes. Agent verification: `./scripts/check-revision-roadmap-alignment.sh` passed; `git diff --check` passed. Go tests/race tests: **not run**. Benchmarks: **not run**. Graphify refresh: **skipped/not run** because no semantic backend was explicitly approved and available; no graph artifacts were changed or staged.
- **Open gaps/blockers:** Targeted prerequisite block remains required before Phase 1. Linux result interpretation needs one later fresh scenario; OS kernel OOM log evidence needs a dedicated Linux VM; Go worker-pool runtime evidence, Kubernetes finalizer semantics, cloud IAM/cost, Git bisection, and container CI smoke testing remain open. The revision-alignment check only guards Phase 1–7 title/order and map coverage; the roadmap remains the authority for all curriculum scope.

## Resume here next session
1. Inspect Git status, latest commit and upstream sync using the start-of-session protocol in `AGENTS.md`.
2. Continue with the Go concurrency prerequisite: define bounded worker-pool shutdown/error semantics, then implement the targeted lab with tests and `go test -race`.
3. Reassess only the worker-pool exit test, then continue the focused prerequisite plan.
4. Update progress records only from observed evidence.

## Next action
Define bounded worker-pool cancellation, channel-closure, and result-ownership semantics before implementation.

## Decision log
- Keep `ROADMAP_2030_REVISED.md` as the curriculum source of truth.
- Keep Go/control-plane/distributed systems as the core path; Rust remains a sequenced follow-up unless explicitly requested earlier.
- Keep phase revision notes and cheatsheets under `revision/`; `revision/ROADMAP_MAP.md` maps them to roadmap sections and defines the agent update contract.
- Enforce that `revision/ROADMAP_MAP.md` and `revision/notes/core-phases.md` retain the roadmap's Phase 1–7 titles/order with `scripts/check-revision-roadmap-alignment.sh`, which the pre-commit hook runs before the optional Graphify step.
