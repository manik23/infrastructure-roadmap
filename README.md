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


## Resume reliably across sessions

Use Git plus `progress/` as the source of truth, not chat memory.

1. At the start of every new session, paste `prompts/01-START-SESSION-RESUME.md` or ask the assistant to follow it. The agent must inspect Git and progress records before resuming.
2. Optional local check: run `./scripts/roadmap-status.sh` from the repository root. This script is read-only; it does not fetch or change anything.
3. If there is an upstream remote and a network connection, the agent can run `git fetch --prune` to refresh remote refs, then compare local HEAD with upstream. If local work is dirty or ahead/behind, the agent must preserve it and ask how to proceed.
4. At the end of a session, paste `prompts/02-END-SESSION-CHECKPOINT.md`. The agent records verified work in `progress/CURRENT_STATE.md`, `weekly-log.md`, `PHASE_TRACKER.md` and `skill-gaps.md` as needed.
5. Review the changes, then commit and push the progress records together with the code/docs/tests they describe. Until then, progress is local and not shared.
6. On the next session/device, repeat the start prompt. If the repo is clean and synced with upstream, the agent resumes from that pushed checkpoint.

### Suggested checkpoint commands (review before executing)

```bash
git status --short --branch
git diff --stat
git diff --cached --stat
git add progress/ path/to/changed-files
git commit -m "learn(phase-1): describe verified checkpoint"
git push
```

Replace `path/to/changed-files` and the message with the actual work. Do not blindly add all files if unrelated changes or secrets may be present. The agent should not run commit/push unless you explicitly authorize it. Never use reset/clean to resolve a checkpoint mismatch.
