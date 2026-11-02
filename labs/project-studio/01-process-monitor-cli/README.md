# 01 — Linux Process-Monitor CLI

**Roadmap:** Phase 1, Week 1
**Status:** Not started
**Primary resource:** `proc(5)` and `proc_pid_status(5)` manual pages —
**required before practice**; they define the Linux data interface.

## Problem and boundary

Build a read-only Go CLI that reports a requested PID's command, state, RSS,
open-FD count, and selected socket information from an injectable `/proc`
root. It runs only in a Linux VM/container. It is not `top`, an eBPF agent, or
a privileged cross-host monitor.

## Learner-owned increments

1. Write parser tests from checked-in fixture text, including a process name
   containing spaces/parentheses where relevant.
2. Implement PID validation and state/RSS parsing; return a useful error when
   the target exits during collection.
3. Add FD enumeration with bounded error handling for permission-denied FDs.
4. Add cancellation-aware repeated sampling and JSON/text output.

**Invariant:** each sample identifies its source and collection time; an
unreadable `/proc` file is not silently reported as zero usage.

**Evidence:** fixture tests plus a disposable Linux-container PID observation.
Record the command and environment; privileged host collection is not needed.
