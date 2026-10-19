# Skill Gaps and Review Queue

Only add a gap when evidence identifies it. Keep the list ordered by dependency and relevance to the active roadmap phase.

| Priority | Domain / concept | Evidence of gap | Planned practice / resource | Reassessment criterion | Status |
|---|---|---|---|---|---|
| P1 | Linux service triage (`ss`, `lsof`, `/proc`, `strace`) | Baseline Linux diagnostic on 2026-10-19: unable to specify process/port, FD, and syscall checks. | Run a short local HTTP-service troubleshooting lab; document commands, observed signals, and interpretation. | Independently propose a safe, ordered diagnostic plan for a listener, FD, syscall, and log failure scenario. | OPEN |
| P1 | HTTPS request path and failure classification | Baseline networking explanation on 2026-10-19 omitted TLS ordering and misclassified HTTP auth/routing errors as timeouts. | Review request phases and capture a safe local `curl -v` trace with labeled outcomes. | Correctly classify DNS/connect/TLS/response failures and HTTP statuses without notes. | OPEN |
| P1 | File-backed memory, page cache, RSS, and OOM diagnosis | Baseline OS response on 2026-10-19 conflated page cache with page tables/virtual memory and did not distinguish RSS from cache. | Complete a small local file-read observation lab and explain cache/reclaim plus OOM evidence. | Accurately explain page cache and RSS, then interpret a kernel OOM record. | OPEN |
| P2 | Go worker-pool lifecycle and error semantics | Baseline review on 2026-10-19 identified cancellation and joining but did not define receive cancellation, error ownership, or provide runnable evidence. | Implement and test a bounded cancellable worker pool. | Race-clean tests verify cancellation, shutdown, and declared error behavior. | OPEN |
| P1 | Kubernetes finalizer/idempotent cleanup semantics | Baseline controller deletion response on 2026-10-19 treated an external `404` as a reason to keep retrying rather than successful absence. | Add a finalizer cleanup test for `404` and transient-failure responses. | Explain and demonstrate finalizer removal after successful or already-absent cleanup. | OPEN |
| P1 | Cloud IAM least privilege and cost safety | Learner stated on 2026-10-19 they need to learn the basic IAM and cost-control scenario. | Complete a provider-specific policy-simulation/design lab and create a sandbox budget/TTL checklist. | Scope storage actions/resources correctly and articulate budget, quota, tagging, and teardown controls. | OPEN |
| P2 | Git regression isolation and safe revert | Baseline workflow response on 2026-10-19 did not identify `git bisect` and proposed reverting all commits after a known-good tag. | Practice a disposable-repository `git bisect run` and targeted revert branch. | Locate a first bad commit and revert only that commit while preserving shared history. | OPEN |
| P2 | Container runtime diagnosis and CI smoke tests | Baseline container response on 2026-10-19 identified a missing copied binary but not common `ENOENT` causes or an automated runtime check. | Build and run a small multi-stage Go image with a health-check CI job. | Diagnose a seeded container startup failure and show a passing build/run/probe CI sequence. | OPEN |

## Rules
- Distinguish `not assessed` from `does not know`.
- Do not infer weakness from a single missed question; probe understanding and practice.
- Close a gap only when the reassessment criterion is demonstrated.
- If a topic is deferred by design, track it in `PHASE_TRACKER.md`, not as a weakness.
