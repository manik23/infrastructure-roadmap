#!/usr/bin/env bash
set -euo pipefail

# Revision documents are evidence/retrieval aids. The roadmap owns phase names
# and ordering, so reject a commit if the derived phase map drifts from it.
repo_root=$(git rev-parse --show-toplevel)
cd "$repo_root"

roadmap='ROADMAP_2030_REVISED.md'
map_file='revision/ROADMAP_MAP.md'
notes_file='revision/notes/core-phases.md'

for required_file in "$roadmap" "$map_file" "$notes_file"; do
  if [[ ! -f "$required_file" ]]; then
    echo "FAILED: required roadmap-alignment file is missing: $required_file" >&2
    exit 1
  fi
done

mapfile -t roadmap_phases < <(rg '^## Phase [1-7] — ' "$roadmap" | sed 's/^## //')
mapfile -t note_phases < <(rg '^## Phase [1-7] — ' "$notes_file" | sed 's/^## //')

if [[ ${#roadmap_phases[@]} -ne 7 ]]; then
  echo "FAILED: expected exactly seven core-phase headings in $roadmap; found ${#roadmap_phases[@]}." >&2
  exit 1
fi

if ! diff -u <(printf '%s\n' "${roadmap_phases[@]}") <(printf '%s\n' "${note_phases[@]}"); then
  echo "FAILED: revision/notes/core-phases.md phase headings or order drift from $roadmap." >&2
  exit 1
fi

for phase in "${roadmap_phases[@]}"; do
  if ! rg -F -q "| $phase |" "$map_file"; then
    echo "FAILED: revision/ROADMAP_MAP.md has no routing row for roadmap heading: $phase" >&2
    exit 1
  fi
done

if [[ $(rg -c '^\| Phase [1-7] — ' "$map_file") -ne 7 ]]; then
  echo "FAILED: revision/ROADMAP_MAP.md must contain exactly seven Phase 1–7 routing rows." >&2
  exit 1
fi

if [[ $(rg -c '^\*\*Source:\*\* \[Roadmap Phase [1-7]\]' "$notes_file") -ne 7 ]]; then
  echo "FAILED: revision/notes/core-phases.md must contain one roadmap source link per core phase." >&2
  exit 1
fi

echo 'PASS: revision phase map and core-phase notes match ROADMAP_2030_REVISED.md headings/order.'
