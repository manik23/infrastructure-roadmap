# Phase Tracker

Update status only when evidence supports it. Suggested statuses: `NOT STARTED`, `IN PROGRESS`, `BLOCKED`, `READY FOR REVIEW`, `PASSED`, `DEFERRED (PLANNED)`. A calendar date alone never makes a phase `PASSED`.

| Phase | Status | Evidence / checkpoint | Remaining gaps | Next action |
|---|---|---|---|---|
| Prerequisite assessment | IN PROGRESS | Initial diagnostics completed 2026-10-19; 2026-10-20 added Linux triage (**Practiced**), HTTPS path (**Demonstrated** for baseline scope), and OS-memory/reclaim (**Practiced**) evidence in `progress/BASELINE_ASSESSMENT.md` | Focused prerequisite exit tests remain: Linux interpretation refresh, OS OOM logs on a Linux VM, Go worker pool, Kubernetes cleanup, cloud IAM/cost, Git bisection, and container CI | Define then implement the bounded Go worker-pool lab with cancellation, shutdown, result/error semantics, and race-clean tests |
| Phase 1 — Systems foundations | NOT STARTED | — | — | Placement decision after assessment |
| Phase 2 — Distributed systems | NOT STARTED | — | — | Follow roadmap sequence |
| Phase 3 — Control planes and scheduling | NOT STARTED | — | — | Follow roadmap sequence |
| Phase 4 — Production platform engineering | NOT STARTED | — | — | Follow roadmap sequence |
| Phase 5 — AI infrastructure | NOT STARTED | — | — | Follow roadmap sequence |
| Follow-up — Rust systems programming | DEFERRED (PLANNED) | Sequence retained in roadmap; no earlier Rust need found during baseline assessment | Core prerequisite block incomplete | Start when the core path has sufficient momentum or a concrete role/project requires it |
| Follow-up — storage internals | DEFERRED (PLANNED) | Sequence retained in roadmap | — | After core path or role-driven need |
| Follow-up — streaming and durable execution | DEFERRED (PLANNED) | Sequence retained in roadmap | — | After core path or role-driven need |
| Follow-up — kernel / high-performance networking | DEFERRED (PLANNED) | Sequence retained in roadmap | — | After core path or role-driven need |
| Follow-up — advanced AI systems | DEFERRED (PLANNED) | Sequence retained in roadmap | — | After AI infrastructure fundamentals |
| Follow-up — formal methods / advanced distributed systems | DEFERRED (PLANNED) | Sequence retained in roadmap | — | Deepen when project demands it |
| Follow-up — security / economics / agentic ops | DEFERRED (PLANNED) | Sequence retained in roadmap | — | Follow after platform foundations |

## Phase pass rule
Mark `PASSED` only when the corresponding roadmap exit criteria are linked to concrete evidence (test results, benchmarks, lab completion, design explanation, or demonstration). Include artifact paths and commit references when available.
