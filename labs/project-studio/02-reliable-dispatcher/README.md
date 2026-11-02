# 02 — Reliable Go Dispatcher

**Roadmap:** Phase 1, Weeks 2–8; capstone Stage A
**Status:** Not started — gated on the prerequisite worker-pool lab
**Primary resource:** Go `context` package and Go diagnostics documentation —
**required before the cancellation and profiling increments**.

## Problem and boundary

Accept jobs through a small API, enforce bounded admission, dispatch work, and
expose observable terminal status. This begins the Infrastructure Lifecycle
Control Plane. Persistence is intentionally basic here; durable lifecycle
semantics arrive in Phase 2.

## Learner-owned increments

1. Define job states, API contract, idempotency boundary, queue capacity, and
   overload response; review tests before implementation.
2. Integrate the verified worker-pool behavior behind the submission API.
3. Add deadlines, graceful shutdown, health/metrics, and structured errors.
4. Produce overload, race, and profile/benchmark evidence under a documented
   workload.

**Invariants:** bounded queued work; one visible outcome per accepted request;
no claim of exactly-once execution; shutdown stops admission and preserves
clear partial-work semantics.

**Evidence:** `go test -race`, overload result, profile explanation, and
before/after benchmark. Start only after the worker-pool review gate passes.
