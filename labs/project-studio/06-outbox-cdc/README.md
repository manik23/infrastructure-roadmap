# 06 — Transactional Outbox and CDC Pipeline

**Roadmap:** Phase 2, Week 18
**Status:** Not started
**Primary resource:** PostgreSQL transaction documentation plus Debezium outbox
documentation — **required before practice**.

## Problem and boundary

Persist a job-state change and outbox row in one PostgreSQL transaction, then
publish through CDC or a polling relay to Kafka/Redpanda. The exercise proves
an at-least-once boundary and consumer deduplication; it does not promise
end-to-end exactly once.

## Learner-owned increments

1. Write a failure table for commit-before-publish, publish-before-ack,
   duplicate delivery, and consumer restart.
2. Write an integration test for atomic state-plus-outbox persistence.
3. Implement a relay/CDC consumer with idempotency keys and replay.
4. Kill/restart a component in a disposable environment and explain observed
   duplicates or recovery.

**Invariant:** a committed state transition has a durable event to publish;
duplicate delivery is tolerated at the consumer boundary.

**Evidence:** integration/failure test output and a documented delivery
guarantee with assumptions.
