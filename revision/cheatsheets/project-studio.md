# Project Studio Cheatsheet

**Source:** [Project studio](../../ROADMAP_2030_REVISED.md#project-studio--progressive-capstone-and-companion-exercises)

Before every increment, write down:

1. Requirement and explicit non-goal.
2. State/ownership invariant and safety versus liveness claim.
3. Failure and resource model: timeout, duplicate, restart, partition, queue,
   memory, or permission boundary as relevant.
4. First deterministic test, model, or design contract.
5. Evidence command/artifact and what it cannot prove.

Quick routing:

- **Phase 1:** `/proc` CLI → bounded dispatcher.
- **Phase 2:** event log → outbox/CDC → rate limiter/cache under failure.
- **Phase 3:** scheduler and controller recovery.
- **Phase 4:** gateway policy and telemetry/log query slice.
- **Follow-up 2:** WAL/recovery/index laboratory only with role/project pull.

Use the capstone for integration; use companion exercises for one mechanism at
a time. A miniature implementation teaches a model, not production-scale
expertise.
