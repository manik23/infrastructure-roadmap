# ADR 0002: One progressive capstone with bounded companion exercises

**Date:** 2026-11-02
**Status:** Accepted

## Context

The roadmap covers rate limiting, caching, event logs, CDC, gateways,
observability, scheduling, and storage internals. Treating each as a full
standalone clone would create many incomplete repositories and obscure the
main target: reliable infrastructure control planes. Omitting implementation
entirely would leave important mechanisms theoretical.

## Decision

Keep the Infrastructure Lifecycle Control Plane as the single progressive
capstone. Add a project-studio series of small, test-first exercises mapped to
the existing roadmap phases and follow-up sequence.

Every exercise must state its requirements, invariants, failure/resource
model, explicit non-goals, learner-owned implementation steps, and evidence
target before code is added. The learner writes or reviews the first test or
design contract before implementation. Mature systems such as PostgreSQL,
Redis, Kafka/Redpanda, OpenTelemetry, and Envoy are used when their real
integration or operational behavior is the learning target.

## Consequences

- The roadmap keeps its Phase 1–7 order, resources, deliverables, and exit
  criteria. Studio work is mapped to those phases; it is not an extra phase.
- “Mini Kafka” and the storage engine are deliberately bounded simulations,
  not claims of broker/database implementation expertise.
- The scheduler and dispatcher extend the capstone; the other exercises
  provide focused mechanisms and may contribute components or evidence.
- A role-gated storage lab remains after the core. Rust remains deferred by
  design unless a concrete role or project justifies earlier work.

## Validation

Review `labs/project-studio/` before beginning each exercise, retain exact
test/benchmark/fault-test output, and run
`./scripts/check-revision-roadmap-alignment.sh` after roadmap or derived-note
changes.
