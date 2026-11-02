# 08 — API Gateway Policy Slice

**Roadmap:** Phase 4; extend in Phase 5 only if the AI gate is met
**Status:** Not started
**Primary resource:** Envoy or chosen gateway's current official docs —
**required before implementation** because configuration/API behavior is
version-sensitive.

## Problem and boundary

Place a narrow policy boundary in front of the platform: request identity,
authentication handoff, deadline propagation, rate policy, correlation IDs,
and safe routing. Use Envoy or a deliberately thin Go component; do not build
a general gateway/proxy.

## Learner-owned increments

1. Write route, identity, timeout, and failure-policy tests using a fake
   upstream.
2. Implement/compose request-ID and deadline propagation.
3. Integrate the rate-limit decision and record reject/timeout metrics.
4. If Phase 5 is activated, add model/tenant budget routing and fallback tests.

**Invariant:** policy is enforced before expensive work; user-visible timeout
semantics and audit identity survive routing.

**Evidence:** integration tests for allow/reject/timeout paths and traces
across gateway to upstream. No inference feature is implied before the AI gate.
