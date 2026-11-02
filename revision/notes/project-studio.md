# Project Studio Revision Notes

**Source:** [Roadmap project studio](../../ROADMAP_2030_REVISED.md#project-studio--progressive-capstone-and-companion-exercises)
**Status:** Not started — catalog/scaffolds exist; no exercise evidence exists.

The studio keeps one progressive control-plane capstone while isolating hard
mechanisms in small exercises. Before coding, each brief in
[`labs/project-studio/`](../../labs/project-studio/README.md) defines the
problem, invariant, expected signal, common limitation, and required evidence.
The learner supplies the logic after a test/model/design-review gate.

| Exercise | Core model / why it matters | Expected signal and limitation |
|---|---|---|
| Process monitor | `/proc` is an observable kernel interface; process diagnostics must distinguish missing permission from zero usage. | Fixture tests plus a disposable Linux observation; not a privileged fleet monitor. |
| Dispatcher | Bounded queue and explicit terminal outcomes prevent overload/leak ambiguity. | Race/overload/profile evidence; not durable or exactly-once yet. |
| Rate limiter | An admission algorithm needs a clock and stated atomicity/failure model. | Deterministic boundary tests; partitions cannot yield unstated global guarantees. |
| Cache | Durable state is truth; TTL/invalidation/stampede handling make cache behavior safe. | Hit/miss/invalidation and overload signals; not a cache-server clone. |
| Mini event log | Ordering is per partition and replay needs an explicit offset/checkpoint. | Ordering/replay tests; no broker replication/rebalance claim. |
| Outbox/CDC | Atomic database state plus outbox narrows the publish gap; consumers tolerate duplicates. | Restart/replay evidence; not end-to-end exactly once. |
| Scheduler | Capacity/quota safety and eventual reconciliation are distinct from fairness trade-offs. | Fixed-snapshot policy comparison and recovery tests; not Kubernetes scheduler replacement. |
| Gateway | Identity, deadlines, policy, and correlation must survive routing. | Allow/reject/timeout integration tests; not a general proxy. |
| Log slice | Structured, redacted correlated telemetry supports diagnosis. | Query/redaction/backpressure evidence; not a log platform. |
| Storage lab | A stated WAL durability boundary permits recovery reasoning. | Crash/replay tests; no production DB/replication claim. |

**Evidence links:** none yet. A scaffold is not implementation, testing, or
demonstration evidence.
