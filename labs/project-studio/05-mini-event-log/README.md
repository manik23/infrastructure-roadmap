# 05 — Mini Partitioned Event Log (“Mini Kafka”)

**Roadmap:** Phase 2, Week 17
**Status:** Not started
**Primary resource:** Kafka documentation on partitions, offsets, and consumer
groups — **required before practice**.

## Problem and boundary

Model an append-only, single-process event log with a fixed partition count,
per-partition offsets, and consumer checkpoints. Its purpose is to make
ordering, replay, and ownership visible before using Kafka or Redpanda. It has
no replication, leader election, storage durability guarantee, or rebalance
protocol.

## Learner-owned increments

1. Write tests defining ordering within a partition and the absence of global
   ordering across partitions.
2. Implement append/read-by-offset with explicit duplicate and invalid-offset
   behavior.
3. Add a consumer checkpoint/replay simulation and a crash/restart test.
4. Compare the missing semantics with an equivalent Kafka/Redpanda exercise.

**Invariant:** offsets are monotonic per partition; a consumer's processed
record boundary is explicit rather than inferred from a successful handler.

**Evidence:** ordering/replay tests and a short comparison of this model's
omitted broker guarantees.
