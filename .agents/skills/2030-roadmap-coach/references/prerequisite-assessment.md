# Prerequisite Assessment Protocol

## Goal
Determine the minimum targeted preparation needed before Phase 1 of `ROADMAP_2030_REVISED.md`. This is a diagnostic, not a beginner course and not a test of personal worth. Skip topics only when the learner can demonstrate them.

## Domains to assess

1. **Linux and debugging:** process/thread basics, file descriptors, signals, `/proc`, `strace`, `lsof`, `ss`, `ip`, logs and simple shell automation.
2. **Networking:** DNS, TCP lifecycle/retransmissions, TLS, HTTP, connection timeouts, L4/L7 balancing and a basic client/server trace.
3. **OS and memory:** processes vs threads, virtual memory, page cache, CPU scheduling, locks, context switching, blocking I/O and latency/throughput distinction.
4. **Go:** interfaces/errors, contexts, goroutines/channels, synchronization, graceful shutdown, tests, `-race`, benchmarks and `pprof` basics.
5. **Kubernetes:** Pods/Deployments/Services, selectors, config/secrets, RBAC, CRDs at a high level, reconciliation and basic `kubectl` debugging.
6. **Cloud fundamentals:** regions/AZs, network boundaries, IAM/least privilege, object/block/file storage, quotas, billing and safe cleanup.
7. **DSA:** hash maps, arrays/lists, stacks/queues, trees/heaps, graphs, sorting/search, complexity analysis and basic dynamic programming.
8. **Engineering workflow:** Git branches/rebase/revert/bisect basics, code review, CI, containers, automated tests and concise design notes.

## Assessment method

- Start with a short self-rating for each domain: comfortable / partly comfortable / rusty / new. Treat it as a hypothesis, not proof.
- For every domain, use one practical question or task and a short explain-back.
- Adapt the difficulty based on the learner's response. Do not force the same test on a senior engineer and a beginner.
- If the learner cannot use the required tools in the current environment, mark it as blocked/not assessed and give a local alternative; don't confuse environment limitations with knowledge gaps.
- Ask one question at a time in interactive tutoring unless the user requests a complete assessment sheet.
- Save evidence, result, gap and recommended remediation in `progress/BASELINE_ASSESSMENT.md`.

## Suggested quick diagnostics

- Linux: identify the process listening on a port and explain how to find its open files; interpret a short `strace` excerpt.
- Networking: narrate DNS → TCP → TLS → HTTP for an HTTPS request and diagnose a timeout versus connection refusal.
- OS/memory: explain when page cache helps, why random I/O differs from sequential I/O, and a likely consequence of unbounded memory growth.
- Go: explain cancellation propagation; write or review a bounded worker pool; identify a race; explain what CPU versus heap profiles show.
- Kubernetes: diagnose a Pod in Pending/CrashLoopBackOff using events, logs, resource requests and configuration; explain desired vs actual state.
- Cloud: sketch a least-privilege service identity for reading one object-storage prefix and explain how to avoid surprise spend.
- DSA: solve a suitable easy/medium problem while explaining complexity and edge cases; scale difficulty to evidence.
- Workflow: describe a safe change workflow from branch to tested reviewable commit and how to locate/regress a suspected bad commit.

## Placement outcomes

- **Ready:** can explain and demonstrate enough to start Phase 1. Keep gaps as just-in-time refreshers.
- **Ready with refreshers:** proceed while scheduling specific 1–2 week refreshers for weak areas.
- **Prerequisite block needed:** do focused work before Phase 1 only for foundational blockers (e.g. cannot yet build/test Go code, use Git or navigate Linux tooling).
- **Not assessed:** there is insufficient evidence or the environment prevents assessment. Do not infer failure or mastery.

The assessment should result in a concise personalized plan with only the necessary prerequisite work, an environment checklist and a date for reassessment. Do not automatically assign all four prerequisite weeks if the learner already demonstrates the skills.
