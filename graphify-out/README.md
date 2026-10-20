# Graphify Outputs

This directory contains reviewed, derived repository-graph artifacts.

Run the guarded refresh only after approving the configured Graphify semantic backend for the session:

```bash
GRAPHIFY_SESSION_GRAPH_APPROVED=1 ./scripts/update-roadmap-graph.sh
```

Successful updates must contain `graph.json` and `GRAPH_REPORT.md`. Review them before staging. If the refresh is skipped or fails, retain the prior valid graph and record the outcome in the session checkpoint; do not claim an update.

Without a configured supported semantic backend, the script prints `SKIPPED` and exits successfully so ordinary commits are not blocked. A configured backend that fails causes the hook to block the commit.
