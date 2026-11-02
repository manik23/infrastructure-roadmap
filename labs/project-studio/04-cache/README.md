# 04 — Shared/Distributed Cache

**Roadmap:** Phase 2, Week 19
**Status:** Not started
**Primary resource:** DDIA caching discussion — **recommended after the first
design attempt**; it helps evaluate invalidation trade-offs.

## Problem and boundary

Add a cache-aside read path for one derived job-status or quota view, first
with a local fake and then an existing shared cache such as Redis. This is not
a Redis implementation or a source of truth.

## Learner-owned increments

1. Write a state table for hit, miss, expired entry, invalidation, stale read,
   and cache outage.
2. Test TTL and explicit invalidation with an injected clock.
3. Add request coalescing/single-flight behavior and test that a miss does not
   trigger an unbounded backend stampede.
4. Measure a controlled hot-key/read workload and report limits.

**Invariant:** durable state remains authoritative; cache failure degrades
performance or availability according to a documented policy, never silently
changes job truth.

**Evidence:** deterministic state tests, stampede test, and a bounded
measurement with cache-hit/miss and latency signals.
