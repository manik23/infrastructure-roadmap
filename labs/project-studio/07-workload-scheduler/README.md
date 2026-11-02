# 07 — Multi-Tenant Workload Scheduler

**Roadmap:** Phase 3, Weeks 26–32; capstone Stage C
**Status:** Not started
**Primary resource:** Kubernetes scheduler framework concepts and the
Kubebuilder Book — **required before controller integration**.

## Problem and boundary

Extend the capstone with workload admission, filters, scores, quotas,
priorities, and fairness across simulated heterogeneous pools. Only after its
simulation is tested does a Kubernetes controller reconcile a real workload.
It is not a replacement Kubernetes scheduler.

## Learner-owned increments

1. Specify workload state transitions, resource model, fairness metric, and
   safety/liveness properties.
2. Write deterministic filter/score and quota tests from a fixed cluster
   snapshot.
3. Compare two policies under the same generated workload; report utilization,
   queue wait, fairness, and starvation limits.
4. Add idempotent reconciliation, finalizer cleanup, restart, and resource-loss
   tests with `envtest`.

**Invariants:** no allocation exceeds capacity or tenant quota; reconciliation
converges after duplicate events/restarts; policy trade-offs are measured.

**Evidence:** scheduler comparison, controller recovery/finalizer tests, and
an ADR for the selected policy.
