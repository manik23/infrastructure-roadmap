# Start Session: Resume from the Last Checkpoint

You are continuing the 2030 infrastructure roadmap. Follow `AGENTS.md` and the `2030-roadmap-coach` skill. Do not start new material until state recovery is complete.

## Required actions
1. Inspect the repository root, active branch, working-tree status, staged/unstaged changes, untracked files, latest commit, and upstream tracking status.
2. If a remote upstream exists and network access is available, run `git fetch --prune` to refresh remote-tracking refs, then compare HEAD with upstream. Fetch is allowed; do not commit, push, reset, clean, stash, rebase or switch branches.
3. Read `progress/CURRENT_STATE.md`, `progress/PHASE_TRACKER.md`, `progress/skill-gaps.md`, the latest relevant entry in `progress/weekly-log.md`, and the relevant section of `ROADMAP_2030_REVISED.md`.
4. Report in this exact structure:
   - Repository / branch
   - Latest commit and upstream sync (`synced`, `ahead`, `behind`, `diverged`, `unknown`, or `no upstream`)
   - Local changes (staged, unstaged, untracked)
   - Last completed and verified checkpoint
   - Current phase / task
   - Evidence already present
   - Open gaps / blockers
   - Recommended next action
5. If there are uncommitted changes or local commits not pushed, preserve them and ask whether to continue from local state or the latest pushed checkpoint. Never discard or silently switch to another commit. If the remote cannot be checked, say sync is unverified.
6. Resume from the latest shared checkpoint only when the working tree is clean and HEAD is confirmed synced with upstream. Progress files must be consistent with the code/test evidence.
7. Continue with one bounded task and follow the skill's full learning workflow.

Do not ask the learner to repeat information already recorded in the repository.
