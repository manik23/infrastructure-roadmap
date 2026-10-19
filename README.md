# 2030 Infrastructure Roadmap Kit

A repository kit for executing the 2030 Distributed Systems, Hyperscale Control Plane and AI Infrastructure roadmap with an AI coding assistant.

## Contents

- `AGENTS.md` — persistent repository instructions and engineering standards.
- `ROADMAP_2030_REVISED.md` — curriculum source of truth.
- `.agents/skills/2030-roadmap-coach/SKILL.md` — reusable skill for study, implementation, verification and phase assessment.
- `.agents/skills/2030-roadmap-coach/references/` — prerequisite diagnostics, learning workflow and evidence rubric.
- `prompts/00-START-PREREQUISITE-ASSESSMENT.md` — start prompt for interactive prerequisite assessment.
- `progress/` — editable baseline, weekly log and skill-gap templates.

## Quick start

1. Extract this kit into the root of your learning repository (or copy its contents into an existing repository after reviewing conflicts).
2. Open the repository in your AI coding assistant.
3. Load/use the `2030-roadmap-coach` skill if your assistant supports Agent Skills; otherwise ask it to read the skill file explicitly.
4. Copy `prompts/00-START-PREREQUISITE-ASSESSMENT.md` into the assistant and follow the interactive assessment one domain at a time.
5. Update the progress templates with observed evidence. Do not mark a skill as mastered simply because you completed a tutorial.
6. Start Phase 1 only after the placement decision and required refreshers are clear.

## Important conventions

- The standard repository instruction filename is `AGENTS.md` (uppercase), not `agent.md`.
- The skill uses the Agent Skills format with YAML frontmatter and is placed in `.agents/skills/` for project-level discovery in compatible tools.
- If your tool expects a different skill directory, move/copy `.agents/skills/2030-roadmap-coach/` to that tool's supported path.
- Keep `ROADMAP_2030_REVISED.md` as the authoritative curriculum; keep progress files factual and editable.
- The prompt asks the assistant to start with a self-rating and then assess one domain at a time, avoiding an overwhelming all-at-once test.
