#!/usr/bin/env bash
set -euo pipefail

# Rebuild the repository graph as a reviewed, derived checkpoint artifact.
# The explicit approval guard prevents an end-session action from silently using
# a configured semantic backend or incurring provider cost.
if [[ "${GRAPHIFY_SESSION_GRAPH_APPROVED:-}" != "1" ]]; then
  echo "SKIPPED: set GRAPHIFY_SESSION_GRAPH_APPROVED=1 after approving the configured Graphify backend."
  exit 0
fi

if ! command -v graphify >/dev/null 2>&1; then
  echo "FAILED: graphify is not installed or not on PATH." >&2
  exit 1
fi

if [[ -z "${GEMINI_API_KEY:-}${GOOGLE_API_KEY:-}${MOONSHOT_API_KEY:-}${ANTHROPIC_API_KEY:-}${OPENAI_API_KEY:-}${DEEPSEEK_API_KEY:-}" ]]; then
  echo "SKIPPED: no supported Graphify semantic backend is configured; graph artifacts were not refreshed."
  exit 0
fi

graphify . --no-viz

if [[ ! -f graphify-out/graph.json || ! -f graphify-out/GRAPH_REPORT.md ]]; then
  echo "FAILED: Graphify exited without both graphify-out/graph.json and graphify-out/GRAPH_REPORT.md." >&2
  exit 1
fi

echo "UPDATED: review graphify-out/graph.json and graphify-out/GRAPH_REPORT.md before staging graph artifacts."
