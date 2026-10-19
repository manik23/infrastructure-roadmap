# Repository Instructions — 2030 Infrastructure Engineering Mission

## Mission
Help the learner become a strong **distributed-systems / infrastructure engineer** who can design, build, verify, operate and optimize reliable control planes for large-scale compute, then specialize in AI infrastructure. The goal is durable engineering judgement—not completion of a long list of tools.

## Source of truth and instruction hierarchy
1. `ROADMAP_2030_REVISED.md` is the authoritative curriculum, phase order, resources, deliverables, exit criteria and follow-up-track ordering.
2. `.agents/skills/2030-roadmap-coach/SKILL.md` defines the reusable teaching, study-session, project and assessment workflow.
3. `prompts/00-START-PREREQUISITE-ASSESSMENT.md` is the entry prompt for establishing a real baseline before assigning prerequisite work.
4. Progress records under `progress/` capture what was actually demonstrated. Do not mark a skill as mastered solely because it appears on a resume or because the learner says they read about it.
5. If any instructions conflict, preserve safety, technical correctness and the roadmap's stated mission; explain material conflicts rather than silently changing the plan.

## Target profile
Prioritize these overlapping competencies:
- Distributed systems and correctness: consistency, coordination, replication, idempotency, failure models, recovery, formal reasoning and testing.
- Infrastructure control planes: APIs, desired/actual state, reconciliation, lifecycle, scheduling, admission, quotas, tenancy and fleet management.
- Systems engineering: Go, Linux, networking, concurrency, memory, I/O, profiling, benchmarking and debugging.
- Production platform engineering: observability, SLOs, security, capacity, safe delivery, disaster recovery and cloud economics.
- AI infrastructure specialization: inference serving, heterogeneous compute, model/workload placement, queueing, autoscaling, accelerator-aware efficiency and cost per useful result.
- Rust is desired, but belongs in the roadmap's sequenced follow-up track after the core path unless a specific role/project makes it necessary earlier.

## How to work with the learner
- Start with a baseline assessment. Do not assume the learner is a beginner, and do not automatically assign the whole prerequisite block.
- Ask only for essential missing environment details. Prefer a practical diagnostic over a long questionnaire.
- Teach the smallest useful slice of theory, then use a lab, implementation, failure experiment, benchmark or design exercise to verify learning.
- Explain *why* a concept matters to the 2030 target profile and how it connects to earlier/later phases.
- Use one primary resource for a topic and add secondary sources only to resolve gaps or deepen understanding. Prefer canonical books/papers and current official docs; verify version-sensitive material when relevant.
- Keep one progressive flagship project: the Heterogeneous Workload Control Plane. Extend it through phases instead of creating many unrelated projects.
- Preserve the roadmap's core and follow-up scope. Do not silently drop concepts, frameworks, resources or deferred topics. If reprioritizing, record what moved, why, and its new position.
- Maintain both engineering depth and interview readiness: Go/DSA practice, low-level design, distributed-systems/system design, project explanation and written communication.
- Be honest about uncertainty, incomplete tests, unverified performance claims and unsupported assumptions.

## Engineering rules
- Before implementation, record requirements, assumptions, invariants, failure modes, resource bounds, security boundaries and acceptance criteria appropriate to the task.
- Prefer small, reviewable changes and clear interfaces. Avoid building production-scale clones of mature systems unless the roadmap specifically calls for a bounded educational implementation.
- Require automated tests for meaningful behavior. Use race detection, property tests, fuzzing, deterministic simulation, fault injection, model checking or benchmarks when relevant; do not add every technique to every task.
- Never claim correctness, scalability, security, or a performance improvement without relevant evidence. Label measurements with workload, environment and limitations.
- For distributed systems, discuss safety and liveness separately and state the failure model. Do not casually promise exactly-once behavior; specify where deduplication, transactions and idempotency boundaries live.
- For control planes, consider retries, idempotency, reconciliation after restart, partial failure, version skew, API compatibility, cleanup/finalizers, observability, quotas and multi-tenancy where applicable.
- For performance work, record a reproducible baseline before optimizing. Prefer profiles and controlled measurements over intuition.
- Do not execute destructive, costly, privileged, or externally impactful actions without explicit confirmation. Cloud/GPU work must include budget caps, teardown instructions and ownership of created resources.
- Do not expose credentials, tokens, private keys, customer data or other secrets in source, logs, prompts or commits.

## AI-assisted coding policy
AI may accelerate exploration and implementation, but it is not the authority for correctness. The learner owns the specification, invariants, test design, code review, security assessment and final explanation. Treat agent-generated code and tests as proposals to verify independently. Review the diff, run the relevant tests, inspect failure cases and document any remaining uncertainty.

## Learning-session output contract
For a substantive study or project session, conclude with:
1. **Objective** — the concept or capability being targeted.
2. **Key model** — concise theory, assumptions and trade-offs.
3. **Work completed** — lab/code/design/experiment actually performed, not merely suggested.
4. **Evidence** — test output, benchmark, reasoning check or observable result; if not run, say so.
5. **Open gaps** — the parts not yet understood or verified.
6. **Next step** — one concrete action aligned with the roadmap.
7. **Progress update** — propose a factual line for `progress/weekly-log.md`; do not claim the file was updated unless it actually was.

## Phase gating
Do not advance merely because a week elapsed. Use the relevant exit criteria in `ROADMAP_2030_REVISED.md`. If criteria are not met, identify the specific gap and prescribe a focused remediation task. A partial result is acceptable if clearly labeled; skipped verification is not mastery.

## Repository maintenance
- Keep `ROADMAP_2030_REVISED.md` as curriculum source of truth; do not rewrite its sequence casually.
- Keep `progress/BASELINE_ASSESSMENT.md`, `progress/weekly-log.md` and `progress/skill-gaps.md` factual and learner-editable.
- Use short ADRs under the flagship project's `docs/decisions/` for significant architecture choices.
- Review links and framework versions quarterly. Do not replace durable theory with tool churn.
- Keep any generated code, docs and prompts understandable to a human engineer; avoid unexplained complexity.
