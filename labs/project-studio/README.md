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
3. Implement one learner-owned increment, then review its invariant and code.
4. The agent adds unit tests and edge cases; generated code is a proposal,
   never evidence by itself.
5. Run the stated verification, record exact output, then update progress and
   the mapped revision note.

## Required task card

Before any catalog exercise or capstone stage begins, the coaching session must
state its exact repository path, concept/problem, target-role relevance,
learner-owned implementation scope, explicit acceptance criteria, verification
command/evidence, and deliberate out-of-scope boundary. A roadmap phase title
or project brief alone is not an assignment or proof of completion.

Use [REVIEW.md](REVIEW.md) for every implementation review. When an exercise
is activated, create its learner-owned source/test scaffold before asking for
logic: package layout, contract types/interfaces, test or fixture/model
skeleton, explicit TODO boundaries, run command, and acceptance criteria.
The learner implements the core logic for the active increment. The agent then
reviews it and generates ordinary unit tests and fixtures for its stated
behavior and edge cases. The learner reviews the invariant, failure model, and
generated assertions rather than manually typing boilerplate tests.
Future-gated exercises may retain templates/non-built starter files until
activation so placeholder behavior cannot make `go test ./...` misleadingly
green.

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
