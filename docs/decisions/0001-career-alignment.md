# ADR 0001: Career-aligned roadmap refinement

**Date:** 2026-11-02  
**Status:** Accepted

## Context

The curriculum was refined using the supplied roadmap2030-refinement-kit,
specifically its career-aligned roadmap and audit reviewed on 2026-10-29. The
refinement identifies distributed cloud infrastructure and storage/workload
control planes as the primary career direction, with AI workload and inference
infrastructure as a complementary, role-evidence-gated specialization.

The supplied job-market examples and survey data are dated inputs, not
independently revalidated market claims in this repository. Their role is to
motivate a repeatable role-research process, not to predict hiring or
compensation.

## Decision

Keep ROADMAP_2030_REVISED.md as the repository's single curriculum authority
and merge the career-alignment decisions into it. Do not create a competing
authoritative roadmap.

Preserve the existing Phase 1 through Phase 7 order, core technical coverage,
prerequisite evidence, and revision-map anchors. Add the following:

1. Storage/workload control planes as the primary career direction.
2. AI infrastructure as a conditional extension after role evidence and core
   control-plane capability.
3. A sustainable six-hour default week, eight-hour ceiling, and three-to-four
   hour busy-week fallback.
4. A 120-day role-validation loop using a scorecard, targeted applications,
   project evidence, and quarterly review.
5. One project storyline: reliable dispatcher, then durable lifecycle service,
   then Heterogeneous Workload Control Plane, then production and optional AI
   extension.
6. Explicit tool-scope limits: one primary tool per need and no custom
   database, full streaming platform, kernel agent, or GPU stack without a
   concrete project or role trigger.

## Consequences

- Existing progress records remain factual historical evidence; no assessment
  status changes because of this planning refinement.
- Rust remains a sequenced follow-up. Storage depth, advanced networking,
  eBPF, streaming, and advanced AI serving remain retained but role-driven.
- The career scorecard and application tracker are evidence tools, not quotas.
- The current worker-pool prerequisite lab remains active and is not promoted
  to Phase 1 completion.

## Validation

The roadmap retains the original Phase 1 through Phase 7 headings and order.
Run scripts/check-revision-roadmap-alignment.sh after this change.
