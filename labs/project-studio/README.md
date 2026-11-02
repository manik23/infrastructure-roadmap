# Project Studio

**Roadmap:** [Project studio](../../ROADMAP_2030_REVISED.md#project-studio--progressive-capstone-and-companion-exercises)
**Status:** Scaffolded only — no studio exercise has learner evidence yet.

This is a sequence of small, enjoyable engineering problems attached to the
roadmap—not a demand to build ten production systems. The progressive
**Infrastructure Lifecycle Control Plane** is the primary project. Companion
exercises isolate mechanisms that make that platform credible.

## How we work an exercise

1. Open only the exercise mapped to the active phase; satisfy its gate first.
2. Read the listed bounded primary resource when it is marked required.
3. Write the first test, model, or design contract. Review it together before
   implementation.
4. Fill in one learner-owned TODO at a time. Generated code is a proposal,
   never evidence by itself.
5. Run the stated verification, record exact output, then update progress and
   the mapped revision note.

Use [REVIEW.md](REVIEW.md) for every implementation review. No Go package is
pre-created merely to contain failing placeholder code; create code and tests
together when an exercise is activated.

The numeric identifiers are stable catalog IDs, not permission to start work
out of sequence. The roadmap phase/week gate below determines the actual order.

| # | Exercise | Phase gate | Brief |
|---:|---|---|---|
| 01 | Linux process-monitor CLI | Phase 1 Week 1 | [brief](01-process-monitor-cli/README.md) |
| 02 | Reliable Go dispatcher | Phase 1 Weeks 2–8 | [brief](02-reliable-dispatcher/README.md) |
| 05 | Mini partitioned event log | Phase 2 Week 17 | [brief](05-mini-event-log/README.md) |
| 06 | Transactional outbox and CDC | Phase 2 Week 18 | [brief](06-outbox-cdc/README.md) |
| 03 | Shared/distributed rate limiter | Phase 2 Week 19 | [brief](03-rate-limiter/README.md) |
| 04 | Shared/distributed cache | Phase 2 Week 19 | [brief](04-cache/README.md) |
| 07 | Multi-tenant job scheduler | Phase 3 Weeks 26–32 | [brief](07-workload-scheduler/README.md) |
| 08 | API gateway policy slice | Phase 4, then Phase 5 Week 51 | [brief](08-api-gateway/README.md) |
| 09 | Log aggregation/query slice | Phase 4 | [brief](09-log-aggregation/README.md) |
| 10 | Storage-engine laboratory | Follow-up 2 | [brief](10-storage-engine/README.md) |

The current prerequisite worker-pool lab remains separate at
[`labs/prerequisites/workerpool/`](../prerequisites/workerpool/README.md).
It must be completed before the dispatcher exercise begins.
