# Core Phase Cheatsheets

These are prompts for recall, not proof of mastery. See the mapped roadmap phase for full scope and resources.

## Phase 1 — Systems fundamentals, Go and performance

- Bound queues and concurrency; state overload behavior and shutdown ownership.
- Propagate deadlines/cancellation; test leaks, races, and duplicate error paths.
- Distinguish throughput from latency; record p50/p95/p99, workload, environment, and baseline.
- Diagnose before optimizing: CPU, heap, contention, and I/O profiles answer different questions.

## Phase 2 — Distributed systems correctness and fault tolerance

- State failure model first; separate safety from liveness.
- Timeouts suspect failure; they do not prove it.
- Define idempotency, duplicate-delivery boundary, ordering, persistence, and recovery.
- Test partitions, retries, crashes, and replay against explicit invariants.

## Phase 3 — Kubernetes control planes, resource management and scheduling

- Reconcile desired toward actual state idempotently; stale events and restarts are normal.
- Treat external `404` during deletion as successful absence when that is the desired end state.
- Use status conditions for observable state; finalizers gate deletion only until cleanup completes.
- Separate admission, placement, fairness, autoscaling, and capacity; each solves a different problem.

## Phase 4 — Production platform engineering, SRE, delivery and security basics

- Trace/metric/log signals must answer an operational question.
- SLOs define user-visible objectives; alerts consume error budget, not raw noise.
- Recovery claims require measured RTO/RPO evidence.
- Identity, policy, provenance, and rollback are delivery requirements, not release afterthoughts.

## Phase 5 — AI compute, inference infrastructure and cost-aware orchestration

- Separate TTFT, inter-token latency, queue wait, throughput, and accelerator utilization.
- Batching/KV cache improve some metrics while risking others; name the workload and quality constraint.
- Compare policies under the same workload; expose cost per useful request/token.
- Simulated placement is useful evidence, not production GPU-scheduling proof.

## Phase 6 — Kernel observability, isolation, security and economics

- Namespaces isolate views; cgroups govern resource accounting/limits; capabilities and seccomp constrain privilege.
- Start eBPF work read-only in an isolated environment.
- Capacity and cost trade against latency, availability, utilization, and interruption risk.

## Phase 7 — Consolidation, interviews and portfolio polish

- Design: requirements/SLOs → estimates → API/data → failures/consistency → bottlenecks → security/cost → trade-offs.
- Explain evidence, assumptions, and what was not tested.
- A smaller reproducible project is stronger than broad unverified claims.
