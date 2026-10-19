#!/usr/bin/env bash
set -uo pipefail

if ! git rev-parse --show-toplevel >/dev/null 2>&1; then
  echo "ERROR: current directory is not inside a Git repository." >&2
  exit 2
fi

echo "Repository: $(git rev-parse --show-toplevel)"
echo "Branch/status:"
git status --short --branch
echo
echo "HEAD:"
git log -1 --format='%h %cs %s'

upstream=$(git rev-parse --abbrev-ref --symbolic-full-name '@{upstream}' 2>/dev/null || true)
if [[ -z "$upstream" ]]; then
  echo "Upstream: none configured (push/sync state cannot be inferred)."
else
  echo "Upstream: $upstream"
  counts=$(git rev-list --left-right --count "HEAD...@{upstream}" 2>/dev/null || true)
  if [[ -n "$counts" ]]; then
    read -r ahead behind <<< "$counts"
    echo "Ahead of upstream: $ahead commit(s)"
    echo "Behind upstream: $behind commit(s)"
  else
    echo "Unable to compare HEAD with upstream."
  fi
fi

echo
echo "Working tree detail:"
git status --short
echo
echo "Progress records:"
for file in progress/CURRENT_STATE.md progress/PHASE_TRACKER.md progress/skill-gaps.md progress/weekly-log.md; do
  if [[ -f "$file" ]]; then echo "OK  $file"; else echo "MISSING  $file"; fi
done

echo
echo "This script is read-only. It does not fetch, commit, push or modify files."
