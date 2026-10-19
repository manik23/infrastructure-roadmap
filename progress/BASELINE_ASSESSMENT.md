# Baseline Assessment

**Started:** 2026-10-19
**Placement:** Prerequisite block needed

| Domain | Self-rating | Status | Evidence observed | Gap | Remediation / next action |
|---|---|---|---|---|---|
| Linux and debugging | 5/10 (learner scale) | Review needed | Could identify `journalctl` as a log source and mentioned `lsof`/`netstat` for a port; did not provide a process/port command, open-FD inspection, or syscall/I/O diagnostic. | Process inspection is being confused with filesystem listing; needs a working Linux incident triage sequence using socket, FD, syscall and service-log tools. | Complete a short safe local triage lab and explain the signals from `ss`, `lsof`, `/proc`, `strace`, and `journalctl`; then repeat a fresh diagnostic. |
| Networking | 6/10 (learner scale) | Review needed | Explained that DNS precedes connection setup, HTTPS commonly uses TCP port 443, and a TCP reset relates to refusal. Did not place TLS before HTTP, distinguish `ECONNREFUSED` from reset of an established connection, or locate timeouts across DNS/connect/TLS/response phases; classified auth/route failures as timeouts. | Needs an accurate request-path model and failure classification for client/service debugging. | Review DNS → TCP → TLS → HTTP and use a local `curl` trace to classify refusal, timeout, TLS failure, and HTTP 401/404 responses. |
| OS, memory and I/O | 6/10 (learner scale) | Review needed | Correctly identified `dmesg` as an OOM evidence source and recognized RAM pressure. Explained page cache as page-table addresses/virtual-memory demand paging, which conflates cached file data with address translation; RSS versus cache explanation was incomplete. | Needs a correct model of file-backed page cache, RSS, reclaim/eviction, and OOM evidence. | Read focused OSTEP virtualization/persistence material; use a local file-read experiment to compare process RSS, system cache, and kernel OOM logs. |
| Go and concurrency | 7/10 (learner scale) | Practiced | Code review on 2026-10-19 correctly identified missing `WaitGroup` lifecycle management, ignored context cancellation, and the need to propagate a deadline into blocking work. | Needs to specify worker-pool shutdown semantics, select on `ctx.Done()` while receiving, and define error/result ownership; no runnable concurrency evidence yet. | Implement a small bounded worker pool with cancellation, `WaitGroup`, per-job deadline propagation, and tests including `go test -race`. |
| Kubernetes and controllers | 7/10 (learner scale) | Review needed | Deletion scenario on 2026-10-19 recognized `deletionTimestamp`, status, a deletion state machine, retries, logs, metrics, and alerts. Treated external `404 Not Found` as a possibly still-provisioning resource and proposed retaining the finalizer/retrying. | Needs idempotent reconciliation/finalizer semantics: an absent external resource is normally cleanup success, so it must not indefinitely block deletion; retry only transient cleanup failures with rate limiting. | Review Kubernetes finalizers and controller-runtime reconciliation; implement an `envtest` or fake-client test proving finalizer removal after a `404` cleanup response and retention after a transient error. |
| Cloud, IAM and cost | 2/10 (learner scale) | Review needed | Learner explicitly stated they need to learn the least-privilege service identity and cloud-cost scenario on 2026-10-19. | Needs IAM resource/action boundaries, workload identity, storage permission scoping, budgets, quotas, tagging, and teardown controls. | Complete a provider-specific IAM and safe-sandbox lab using policy simulation or documentation only first; no paid resources required. |
| DSA and complexity | 5/10 (learner scale) | Practiced | On 2026-10-19, proposed the correct one-pass hash-set algorithm for first repeated job ID and stated expected O(n) time/O(n) space. Also selected a max-heap for highest-priority job selection. Edge cases, stable tie-breaking, stated heap complexities, and broader tree/graph evidence remain unassessed. | No demonstrated gap in basic hash-map or max-heap selection; needs to state priority-queue costs and model a stable tie-breaker before readiness is established. | Complete one scheduler-queue exercise with a composite priority key and operation complexities; retain only any demonstrated gap for a targeted DSA refresher. |
| Git, containers, testing and workflow | Not self-rated | Review needed | On 2026-10-19, identified `git log` and a new revert commit as safer than rewriting shared history. Proposed `git rev-parse` and reverting all commits after a good tag to locate the regression; did not identify bisection or a test-driven verification workflow. For a container runtime `ENOENT`, identified a missing copy but not runtime path, executable, dynamic-linker, or architecture diagnosis; CI/test evidence remains unassessed. | Needs `git bisect` workflow, container runtime inspection, and test/CI verification discipline. | Run a disposable-repository `git bisect run` exercise, create a targeted revert branch, and build/run a small container health check in CI. |

## Environment readiness
- [x] Linux host or VM for Linux-specific labs — Docker Desktop verified on 2026-10-19 with an Alpine Linux/arm64 container; use a disposable Linux container first.
- [x] Go toolchain — verified 2026-10-19: Go 1.27.1 on darwin/arm64.
- [x] Git — verified 2026-10-19: Git 2.54.0 (Apple Git-157); use a disposable repository for the bisection lab.
- [x] Container runtime — verified 2026-10-19: Docker Desktop Engine 29.8.2, linux/arm64.
- [ ] `kind` or `k3d` — install and verify only before the Kubernetes finalizer lab.
- [ ] Automated test command — verify `go test ./...`; add `go test -race ./...` for the worker-pool lab.
- [ ] Cloud guardrails — before any cloud experiment, set a small budget alert, quota, expiry tags/TTL, ownership tag, and a documented teardown command. Start with policy simulation only.
- [ ] GPU — not needed for prerequisites; use CPU/simulation. Before any later GPU use, set a hard duration/budget cap and teardown owner.

## Placement decision
Choose one after sufficient evidence:
- [ ] Ready
- [ ] Ready with refreshers
- [x] Prerequisite block needed
- [ ] Not assessed

Foundational Linux, networking, OS-memory, cloud/IAM, and workflow gaps would prevent reliable Phase 1 work. Complete the focused labs below, then reassess their exit tests. No Rust prerequisite was identified; Rust remains a deferred follow-up track.

## Personalized prerequisite plan
Fill only for required gaps. Do not copy the full curriculum here.

| Gap | Primary resource | Estimated effort | Practical deliverable | Exit test |
|---|---|---:|---|---|
| Linux incident triage | MIT Missing Semester: command-line environment and shell tools sections | 2–3 hours | A troubleshooting note for a local HTTP service: listener/PID, open FDs, one `strace` observation, and relevant logs. | Given a slow service scenario, select safe commands and distinguish listener, FD exhaustion, blocking I/O, and failing syscalls. |
| HTTPS request path and client failures | High Performance Browser Networking: DNS, TCP, TLS, and HTTP chapters | 2–3 hours | A short request-timeline note plus `curl -v` output from safe local endpoints, labeling each phase and failure class. | Correctly explain DNS → TCP handshake → TLS handshake → encrypted HTTP, and classify connection refusal, timeouts, TLS errors, and HTTP status responses. |
| Page cache, RSS, and OOM triage | OSTEP: virtualization and persistence chapters | 2–3 hours | A short local file-read observation note that separates per-process RSS from system file cache and cites kernel OOM evidence. | Explain file-backed cache/reclaim accurately and identify OOM evidence with `journalctl -k` or permitted `dmesg` output. |
| Go worker-pool cancellation and shutdown | Go `context` package documentation and Go diagnostics guides | 2–3 hours | A bounded worker pool with a defined shutdown/error policy and cancellation test. | `go test -race ./...` passes; a test proves workers exit on cancellation and blocking work receives a deadline. |
| Kubernetes finalizers and idempotent cleanup | Kubebuilder Book: finalizers and reconciliation chapters | 2–3 hours | A controller test where external cleanup returns `404` and the finalizer is removed; a transient failure retains the finalizer and requeues. | Explain why delete is idempotent and demonstrate correct finalizer behavior in a test. |
| Cloud IAM and cost guardrails | AWS IAM User Guide (roles, identity policies, policy simulator) and AWS Budgets documentation; use as the reference model unless a target employer requires another provider | 4–6 hours | A least-privilege workload-identity policy for two storage prefixes plus a sandbox budget/TTL teardown checklist. | Explain the policy resource/action boundary and show a dry-run or policy-simulator result; name budget, quota, and teardown safeguards. |
| Git regression isolation and safe revert | Pro Git: debugging with Git and undoing changes | 1–2 hours | A disposable-repository `git bisect run` transcript and a revert branch with a targeted regression test. | Locate a seeded bad commit with `git bisect`, revert only that commit on a branch, and show the relevant test passes. |
| Container runtime diagnosis and CI smoke test | Docker Get Started: build and run a containerized application | 2–3 hours | A multi-stage Go image plus CI job that builds it, runs it as non-root, probes a health endpoint, and captures logs on failure. | Diagnose a seeded wrong-path, missing-executable, or dynamic-linker runtime failure and show the CI smoke test passing. |

## Reassessment notes
- 2026-10-19: Baseline diagnostic completed across all eight domains. Reassess each listed exit test before beginning Phase 1.
