# 10 — Storage-Engine Laboratory

**Roadmap:** Follow-up 2; Phase 2 Week 12 supplies concepts
**Status:** Deferred by design until role/project evidence selects it
**Primary resource:** *Database Internals*, WAL/recovery chapters — **required
before practice**.

## Problem and boundary

Implement a deliberately small persistent key/value model with an append-only
WAL, crash recovery, and exactly one index choice (B-tree *or* LSM-like
memtable/SSTable model). It is a laboratory for durability/recovery reasoning,
not the capstone's production persistence layer or a general database.

## Learner-owned increments

1. Write the crash model and WAL record/checkpoint format; test torn/partial
   records and recovery decisions.
2. Implement append and replay before adding an index.
3. Add one index and test read/write behavior plus duplicate/tombstone policy.
4. Benchmark a stated read/write mix and explain recovery, amplification, and
   tail-latency limits.

**Invariant:** after any modeled crash, recovered state reflects a defined
durability boundary; corrupted/torn input is rejected or handled explicitly.

**Evidence:** deterministic crash/replay tests and a reproducible benchmark.
No replication, SQL, or production durability claim is required.
