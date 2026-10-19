# End Session: Record a Reliable Checkpoint

Before ending this roadmap session, follow `AGENTS.md` and the `2030-roadmap-coach` skill.

1. Summarize what was actually changed or learned. Do not describe planned work as completed.
2. Record exact commands run and test/benchmark outcomes. Mark anything not run as `not run`; don't infer success.
3. Update `progress/CURRENT_STATE.md`, append an entry to `progress/weekly-log.md`, and update `progress/PHASE_TRACKER.md` and `progress/skill-gaps.md` if the evidence changes them.
4. Include topic status (`Demonstrated`, `Practiced`, `Review needed`, `Not assessed`, or `Deferred by design`), what remains unverified, and one next action.
5. List changed files and check `git status --short --branch`. Do not commit or push unless the learner explicitly asks in this conversation.
6. Suggest a commit message that includes both implementation and progress records, for example: `learn(phase-1): add bounded worker pool and record verification`.
7. State clearly: `LOCAL / NOT YET SHARED` until the learner commits and pushes. Do not claim persistence to another session/device until a later start-session check confirms the progress commit is on the upstream branch.
8. If there are untracked or unrelated user changes, preserve them and call them out. Never run destructive Git commands to make the tree clean.
