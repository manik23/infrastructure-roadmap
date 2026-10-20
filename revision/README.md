# Revision Layer

This folder is a retrieval aid for active study. [ROADMAP_2030_REVISED.md](../ROADMAP_2030_REVISED.md) remains the curriculum source of truth: phase order, resources, deliverables, and exit criteria belong there.

## Layout

- [ROADMAP_MAP.md](ROADMAP_MAP.md) maps every revision section to its authoritative roadmap heading.
- `notes/` holds learner-specific explanations, misconceptions, evidence links, and open questions.
- `cheatsheets/` holds compact recall prompts, invariants, failure modes, and verification signals.

## Update contract for agents

1. Read the mapped roadmap section and relevant `progress/` records before editing.
2. Add only demonstrated evidence, observed misconceptions, or concise learner-authored explanations to notes. Never mark a topic complete from reading alone.
3. Keep cheatsheets compact. They may summarize stable concepts but must not add, remove, or reorder roadmap scope.
4. Link implementation, test, benchmark, ADR, or postmortem evidence when it exists. State `not run` when verification is absent.
5. Update the matching `progress/` record when evidence changes a status or gap; do not make a cheatsheet claim substitute for progress evidence.
6. Keep Rust in the mapped follow-up section unless a concrete project or demonstrated core-path gap justifies earlier work.
7. When an approved end-session Graphify update succeeds, review its derived artifacts before staging; the graph complements this map but does not replace the roadmap.

## Status labels

Use `Not started`, `In progress`, `Practiced`, `Demonstrated`, `Review needed`, or `Deferred by design`. These are evidence labels, not calendar milestones.
