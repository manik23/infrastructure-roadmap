# 09 — Log Aggregation and Query Slice

**Roadmap:** Phase 4
**Status:** Not started
**Primary resource:** OpenTelemetry logging and trace documentation —
**required before schema/instrumentation work**.

## Problem and boundary

Emit structured, tenant-safe events from the capstone, carry trace/request/job
correlation, and query them through an existing collector/store. The learning
target is observability data quality, backpressure, retention, and redaction—
not an Elasticsearch/Loki replacement.

## Learner-owned increments

1. Define an event schema and a redaction/retention policy; write schema tests.
2. Instrument a request → dispatch → terminal-state path with correlation IDs.
3. Route data through an OpenTelemetry Collector and query one incident-like
   scenario.
4. Apply a bounded queue/drop policy and demonstrate its behavior under load.

**Invariant:** logs never contain secrets or raw sensitive payloads by default;
correlation is sufficient to diagnose a request without pretending logs are a
source of truth.

**Evidence:** a trace/log query, redaction test, and a documented overload
policy with observed signals.
