# 03 — Shared/Distributed Rate Limiter

**Roadmap:** Phase 2, Week 19
**Status:** Not started
**Primary resource:** DDIA, bounded sections on time and coordination —
**required before the shared-store increment**.

## Problem and boundary

Make admission decisions for a tenant/key with a stated burst and sustained
rate. Start with a deterministic clock and a token-bucket or GCRA model, then
compare an in-memory instance with an atomic shared-store implementation. This
is not a globally available limiter across partitions.

## Learner-owned increments

1. Write tests for refill, burst, boundary time, and rejection metadata using
   an injected clock.
2. Implement the local algorithm and explain its space/time bounds.
3. Define the shared-store atomic operation and test two simulated replicas.
4. Inject store timeout/unavailability and choose fail-open or fail-closed per
   endpoint; document the trade-off.

**Invariant:** accepted requests never exceed the declared local/shared model
under its assumptions. Partition behavior is explicitly a policy decision.

**Evidence:** deterministic tests, a two-replica experiment, and the stated
failure model—not a claim of universal global correctness.
