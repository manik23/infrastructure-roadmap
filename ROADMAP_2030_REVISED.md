# 2030 Hyperscale Infrastructure & AI Platform Engineering Roadmap

**Revised:** 19 October 2026  
**Mission:** Become a strong distributed-systems / platform engineer who can design, build, operate and optimize control planes for large-scale compute—and then specialize in AI infrastructure.  
**Core duration:** 60 weeks, approximately 14–15 hours/week.  
**Follow-up tracks:** Rust systems programming plus targeted specializations, taken in dependency order after the core rather than mixed into every phase.

> **North star:** Build systems that turn unreliable, heterogeneous, expensive resources into reliable services. Learn the theory behind correctness and performance; use frameworks to apply it; prove the result with tests, experiments and operational evidence.

This is a career-development plan, not a prediction that any specific framework will dominate in 2030. The durable assets are distributed-systems reasoning, Linux and networking fundamentals, control-plane design, performance analysis, reliability, security, resource scheduling and cost-aware engineering. Revisit tool choices quarterly.

---

## 1. Target profile and scope

### Target roles

Prioritize roles whose work includes several of the following:

- Distributed Systems Engineer / Control Plane Engineer
- Compute Platform / Cluster Management Engineer
- Kubernetes Infrastructure / Scheduling Engineer
- AI Infrastructure / Inference Platform Engineer
- Fleet Management / Capacity Engineering
- Infrastructure Performance / Reliability Engineer

### Core specialization

**Distributed control planes and workload orchestration**, with **AI compute and inference infrastructure** as the growth specialization.

### The stack you should understand

1. **Hardware and host:** CPU, memory hierarchy, storage, accelerators, NUMA, kernel, containers and isolation.
2. **Data plane:** network, RPC, storage I/O, proxies, inference execution and workload processes.
3. **Control plane:** APIs, desired/actual state, reconciliation, scheduling, admission, quotas, policy, lifecycle and recovery.
4. **Platform:** multi-tenancy, developer workflows, deployment, observability, reliability and security.
5. **Economics and product constraints:** latency, throughput, utilization, availability, cost per workload and capacity headroom.

### What this plan does not require initially

You do **not** need to become a chip designer, ML researcher, deep-learning scientist or BGP specialist. Learn enough about these domains to reason about their effect on infrastructure. Go deeper only if your eventual role calls for it.

### Core versus follow-up

- **Core:** Understand the concepts and implement enough to demonstrate them. These are prerequisites for the target role.
- **Follow-up:** Important specialties that are deliberately sequenced after the core. They are not forgotten or silently removed.
- **Role-dependent:** Learn when a target role or team needs them; do not study every tool just because it exists.

---

## 2. Weekly operating model

### Sustainable time budget

Target **14–15 hours/week** while employed:

| Slot | Time | Purpose |
|---|---:|---|
| Mon–Thu | 1.5–2 hours/day | Reading, focused labs, coding or DSA |
| Friday | 1–2 hours | Review, notes, paper summary, design exercise |
| Weekend | 4–5 hours | Project integration, tests, benchmark or failure drill |

If only 8–10 hours are available in a busy week, keep the project and interview tracks alive and stretch the calendar. Do not skip tests to meet a date.

### Rules that keep the roadmap effective

1. **Project-led, not tool-led.** Every tool must support a real requirement in the flagship project or a focused lab.
2. **One primary implementation language initially: Go.** Keep Rust in the follow-up track, as requested.
3. **Theory before complexity.** State the invariant, failure model and success metric before coding.
4. **Measure, don't guess.** Include a benchmark, fault test or operational metric in each meaningful deliverable.
5. **Write down decisions.** For every capstone, produce a short design doc or ADR with alternatives and trade-offs.
6. **Use AI coding tools as assistants, not authorities.** You own requirements, threat models, test design, review and correctness claims.
7. **Review the plan quarterly.** Keep fundamentals stable; update framework choices based on current docs and target job descriptions.
8. **Do not turn every topic into a custom implementation.** Read production designs and build a small model or lab when that is enough to understand the concept.

---

## 3. Entry assessment and lab environment

Skip any baseline topic you can already demonstrate. Your existing Go and Kubernetes experience means this should be a skills assessment, not automatically four weeks of remedial work.

### Baseline checklist

- [ ] **Linux:** process/thread model, signals, file descriptors, `/proc`, permissions, `strace`, `lsof`, `ss`, `ip`, `journalctl`.
- [ ] **Networking:** DNS, TCP lifecycle, retransmissions, flow/congestion control basics, TLS, HTTP, L4/L7 load balancing.
- [ ] **OS and memory:** virtual memory, page cache, allocation, CPU scheduling, locks, context switching, file I/O.
- [ ] **Go:** interfaces, errors, contexts, cancellation, goroutines, channels, `sync`, testing, race detector, benchmarks and `pprof`.
- [ ] **Kubernetes:** Pods, Deployments, Services, RBAC, CRDs, reconciliation and basic `kubectl` debugging.
- [ ] **Cloud:** regions/AZs, VPC/networking, IAM, object/block/file storage, quotas and billing concepts.
- [ ] **DSA:** hash maps, stacks/queues, trees/heaps, graphs, sorting/searching, complexity, basic dynamic programming.
- [ ] **Engineering workflow:** Git, code review, CI, containers, test automation and basic design docs.

If any item is weak, use the corresponding prerequisite resource below and spend 1–2 weeks on it before or alongside Phase 1.

### Lab setup

- Use a **Linux environment** for kernel, eBPF, traffic shaping and low-level profiling. On macOS, use a Linux VM such as Lima, OrbStack or Multipass; verify kernel/tool compatibility for each lab.
- Use `kind` or `k3d` for local Kubernetes. Keep cloud experiments behind budget alerts and teardown scripts.
- Run distributed tests locally first. Use small clusters and simulated resources before paying for a GPU.
- For AI inference, start with a tiny CPU model or an inexpensive short-lived GPU instance. Track and stop all resources after experiments.
- Keep one Git repository for the flagship project, with a `docs/`, `design/`, `benchmarks/`, `runbooks/` and `experiments/` structure.

### Prerequisite resources

| Topic | Main resource | Reference / lab |
|---|---|---|
| Linux tools | [MIT Missing Semester](https://missing.csail.mit.edu/) | [Julia Evans zines](https://wizardzines.com/), [Linux performance tools](https://www.brendangregg.com/linuxperf.html), *How Linux Works* |
| OS concepts | [Operating Systems: Three Easy Pieces](https://pages.cs.wisc.edu/~remzi/OSTEP/) | *Computer Systems: A Programmer's Perspective*; *The Linux Programming Interface* |
| Networking | [High Performance Browser Networking](https://hpbn.co/) | [Beej's Guide](https://beej.us/guide/bgnet/), [TLS 1.3 explained](https://tls13.xargs.org/), [QUIC explained](https://quic.xargs.org/) |
| Go | [A Tour of Go](https://go.dev/tour/) | [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests), [Effective Go](https://go.dev/doc/effective_go) |
| Git | [Pro Git](https://git-scm.com/book/en/v2) | [Learn Git Branching](https://learngitbranching.js.org/) |
| Containers | [Docker Get Started](https://docs.docker.com/get-started/) | [iximiuz Labs](https://labs.iximiuz.com/) |
| Cloud basics | Your provider's IAM, networking and storage docs | [AWS Well-Architected](https://docs.aws.amazon.com/wellarchitected/latest/framework/welcome.html) as a general design reference |
| DSA | [NeetCode roadmap](https://neetcode.io/roadmap) | [Tech Interview Handbook](https://www.techinterviewhandbook.org/) |

---

# PART I — CORE ROADMAP (60 WEEKS)

## Phase 1 — Systems fundamentals, Go and performance

**Weeks 1–8**  
**Outcome:** You can write concurrent services, understand their interaction with Linux and the network, and explain performance results using evidence.

### Concepts and theory

**Operating systems and computer architecture**
- User/kernel boundary, syscalls, processes versus threads, scheduling and context switches.
- Virtual memory, page tables/TLB, page cache, allocation, stack/heap, memory pressure and OOM behavior.
- CPU caches, locality, cache lines, false sharing, branch prediction at a conceptual level, NUMA awareness.
- File descriptors, blocking/non-blocking I/O, `epoll`, `io_uring` concepts, buffered versus direct I/O and zero-copy trade-offs.
- Locks, condition variables, atomics, memory ordering and the difference between safety and liveness.

**Go and concurrency**
- G-M-P scheduler, goroutine lifecycle, channel and `select` behavior, context cancellation and deadline propagation.
- Mutex/RWMutex, atomics, `sync.Once`, `WaitGroup`, `sync.Pool` trade-offs, map safety and data races.
- Worker pools, bounded queues, backpressure, cancellation, graceful shutdown, leak detection and overload behavior.
- Allocation profiles, escape analysis at a useful level, GC trade-offs, heap/CPU profiles and benchmark methodology.

**Networking and service interfaces**
- DNS, TCP handshake/teardown, congestion and flow control, retransmission, keepalive and connection pooling.
- HTTP/1.1, HTTP/2 and HTTP/3/QUIC at a design level; TLS certificates and handshake; L4 versus L7 load balancing.
- API semantics, timeouts, retries, idempotency, pagination and backward-compatible schemas.
- gRPC/Protobuf: deadlines, status codes, streaming, schema evolution and compatibility checks.

**Performance method**
- Throughput versus latency; p50/p95/p99; closed-loop versus open-loop load generation; coordinated omission.
- CPU, allocation, contention and I/O profiling; flamegraphs; benchmark noise and statistically credible comparisons.

### Weekly sequence

| Week | Focus | Deliverable |
|---|---|---|
| 1 | Linux process, memory, file and network inspection; latency and resource basics | Troubleshooting notebook with `strace`, `ss`, `lsof` and `/proc` examples |
| 2 | Go concurrency, memory model, cancellation, race detector | Bounded worker pool with shutdown, cancellation and race tests |
| 3 | Queueing, backpressure, overload and retry behavior | Load generator showing why an unbounded queue causes latency/memory growth |
| 4 | TCP/TLS/HTTP, timeouts and API semantics | Small Go HTTP/gRPC service with deadlines, graceful shutdown and structured errors |
| 5 | Protobuf/gRPC streaming and compatibility | Versioned API with `buf` lint/breaking-change checks and a streaming endpoint |
| 6 | `pprof`, benchmarks, GC, allocations, `perf`/flamegraphs | Baseline performance report plus one evidence-based optimization |
| 7 | CPU cache, synchronization, I/O paths, profiling under load | Compare two implementations; record throughput, p99, allocations and CPU per request |
| 8 | **Buffer and capstone review** | A reliable Go job-dispatch service, design note, reproducible benchmark and failure tests |

### Phase 1 project: reliable job dispatcher

Build a Go service that accepts jobs, enqueues work in a **bounded** queue, dispatches to workers and exposes status. Add request IDs/idempotency, deadlines, graceful termination, health endpoints, metrics and load tests. At this stage, persistence can be basic; durable distributed workflow behavior comes later.

**Exit criteria**
- Explain an actual CPU or allocation profile produced by your code.
- Demonstrate bounded resource usage under overload.
- Pass `go test -race ./...` and meaningful concurrency tests.
- Show a before/after benchmark with test environment, workload and limitations documented.

### Resources

- [Go diagnostics](https://go.dev/doc/diagnostics), [Go memory model](https://go.dev/ref/mem), [Go blog](https://go.dev/blog/)
- [OSTEP](https://pages.cs.wisc.edu/~remzi/OSTEP/), especially concurrency, virtualization and persistence
- [Brendan Gregg's performance resources](https://www.brendangregg.com/)
- *Systems Performance* (Brendan Gregg)
- *Concurrency in Go* (Katherine Cox-Buday)
- [gRPC concepts](https://grpc.io/docs/what-is-grpc/core-concepts/), [Protocol Buffers](https://protobuf.dev/)

---

## Phase 2 — Distributed systems correctness and fault tolerance

**Weeks 9–20**  
**Outcome:** You can reason about partial failure, consistency, replication, durable work and overload—and test correctness instead of relying on happy-path demos.

### Concepts and theory

**Models and failure**
- Processes, messages, network partitions, crash-stop and crash-recovery, omission failures, untrusted timing and partial failure.
- Safety versus liveness; failure detectors; timeouts are suspicions, not proof of failure.
- CAP and PACELC as reasoning tools, not simplistic database labels; FLP intuition and assumptions behind consensus.
- Clock drift, wall clocks versus monotonic clocks, Lamport clocks, vector clocks and hybrid logical clocks (conceptual plus examples).

**Consistency, replication and coordination**
- Linearizability, sequential consistency, eventual consistency, causal ordering and read-your-writes.
- State-machine replication; Raft in depth, Paxos conceptually; leaders, terms/epochs, log matching, quorum, commit indexes and membership changes.
- Leases and fencing tokens, leader election, idempotent writes, duplicate delivery and split brain.
- Replication lag, quorum reads/writes, snapshots, log compaction and recovery.
- CRDTs and gossip/membership (SWIM-style) as alternatives for selected problems; know their trade-offs.

**Transactions and storage fundamentals**
- ACID, isolation levels, dirty/non-repeatable/phantom reads, MVCC, optimistic versus pessimistic concurrency control.
- WAL, checkpoints, B-trees versus LSM trees, compaction, indexes, cache/page behavior and durability boundaries.
- Primary/replica, synchronous/asynchronous replication, partitioning, consistent hashing, rebalancing, resharding and multi-region trade-offs.
- PostgreSQL fundamentals: indexes, `EXPLAIN ANALYZE`, connection pools, vacuum and transaction isolation.

**Messaging, workflows and data flow**
- Delivery semantics: at-most-once, at-least-once and the limits/meaning of “exactly once.”
- Idempotency keys, transactional outbox/inbox, CDC, ordering by partition/key, replay and consumer offset management.
- Event time versus processing time, watermarks, stateful stream processing concepts.
- Sagas, compensating actions, durable workflow orchestration and timeout/retry policies.
- Caching, TTL/invalidation, stampedes, hot keys and stale data.

**Resilience and queueing**
- Deadlines and timeout budgets, exponential backoff with jitter, retry budgets, circuit breakers, bulkheads, admission control, load shedding and graceful degradation.
- Backpressure, Little’s Law, queue growth, head-of-line blocking, tail latency, hedged requests and overload/metastable failure.
- Rate limiting: token bucket and GCRA concepts; avoid retry amplification.

**Correctness engineering**
- Unit/integration/property-based testing, fuzzing, race detection and invariants.
- TLA+ specification and model checking for a small protocol; deterministic simulation; history-based linearizability checking.
- Chaos experiments with network delay/partition, process kill, disk errors and clock skew; distinguish a useful fault test from random disruption.

### Weekly sequence

| Week | Focus | Deliverable |
|---|---|---|
| 9 | Failure models, safety/liveness, CAP/PACELC and FLP intuition | Failure-model table for the job dispatcher |
| 10 | Time, ordering, idempotency, retries and fencing | Duplicate/reordered request tests and a lease/fencing-token design |
| 11 | Consistency models, transactions, MVCC and isolation | Write anomaly examples and tests against PostgreSQL |
| 12 | WAL, B-tree/LSM, replication and durability | Design note comparing two storage approaches; small WAL or log-replay lab |
| 13 | Raft: election, terms, log matching and safety | Complete the early labs from MIT 6.5840 or a comparable Raft lab |
| 14 | Commit, snapshots, membership changes and recovery | Failure/restart experiment with a replicated-state lab |
| 15 | TLA+ and invariants | Model a small leader/lock or job-lease protocol and model-check it |
| 16 | Partitioning, replication, sharding and geo trade-offs | Design a multi-region job metadata service; state latency/consistency choices |
| 17 | Event logs, consumer offsets, ordering, replay and CDC | Idempotent producer/consumer lab using Kafka or Redpanda |
| 18 | Outbox, event time, windows, workflow/saga concepts | Postgres outbox → message stream prototype; document delivery guarantees |
| 19 | Caching, queueing, backpressure, rate limits and overload | Load test illustrating queue growth, retry storm and recovery after load shedding |
| 20 | **Buffer and fault-injection review** | Fault matrix, seeded chaos tests, property/fuzz tests and a postmortem |

### Phase 2 project: make job state durable and correct

Extend the job dispatcher into an asynchronous job-control service with persistent state, an idempotent submission API, retry policies, worker leases/fencing tokens, durable status transitions, a transactional outbox and event consumers. Use PostgreSQL as the durable store rather than building a database from scratch. Run a **separate bounded Raft lab** to learn consensus; do not let a custom database implementation delay the control-plane track.

**Exit criteria**
- State the consistency and failure guarantees of each API.
- Demonstrate safe duplicate requests and recovery after worker/coordinator failure.
- Have at least one protocol model or invariant test, one fault-injection test and a written postmortem.
- Explain when the design can lose work, duplicate work or return stale state—and how it handles each case.

### Resources

- *Designing Data-Intensive Applications* (DDIA), 2nd edition if available to you; use the edition you can access and verify errata/contents.
- [MIT 6.5840 Distributed Systems](https://pdos.csail.mit.edu/6.824/), [Raft paper](https://raft.github.io/raft.pdf), [Raft visualization](https://raft.github.io/)
- [Fly.io Gossip Glomers](https://fly.io/dist-sys/), [TLA+ resources](https://lamport.azurewebsites.net/tla/tla.html), *Practical TLA+* (Hillel Wayne)
- [Porcupine linearizability checker](https://github.com/anishathalye/porcupine), [Jepsen analyses](https://jepsen.io/analyses), [Toxiproxy](https://github.com/Shopify/toxiproxy)
- *Release It!* (Michael Nygard), *Database Internals* (Alex Petrov)
- [PostgreSQL documentation](https://www.postgresql.org/docs/), [Kafka documentation](https://kafka.apache.org/documentation/), [Temporal documentation](https://docs.temporal.io/)

---

## Phase 3 — Kubernetes control planes, resource management and scheduling

**Weeks 21–32**  
**Outcome:** You can design a robust operator/control plane and reason about scheduling, multi-tenancy and heterogeneous resources. This is the most career-critical phase.

### Kubernetes architecture and control-plane concepts

- API server, etcd, scheduler, controller manager, kubelet, CRI, CNI and CSI responsibilities.
- Kubernetes API semantics: `spec` versus `status`, conditions, generations, resource versions, watches, caches and optimistic concurrency.
- Reconciliation loops, idempotency, eventual convergence, rate-limited retries, work queues and resync behavior.
- CRDs, schema validation/defaulting, conversion/versioning, status subresource, finalizers, owner references, garbage collection and deletion lifecycle.
- Leader election, controller concurrency, work partitioning, API throttling and handling stale cache reads.
- Admission, RBAC, service accounts, namespace boundaries, resource requests/limits, quotas and LimitRanges.
- Operator testing: unit tests, fake clients versus `envtest`, integration tests, upgrade/rollback and failure tests.
- Cluster lifecycle and multi-cluster concepts; fleet drift, policy propagation, rollout waves and blast-radius control.

### Scheduling and resource management

- Scheduler pipeline: queueing, filtering/feasibility, scoring, reserve/permit/pre-bind/bind/post-bind concepts.
- Bin-packing versus spreading; affinity/anti-affinity, topology spread, locality, priorities, preemption and disruption budgets.
- Admission control versus placement; queueing delay versus execution time; workload classes and SLO-aware policies.
- Fair-share, quotas, borrowing, starvation, backfilling, gang/co-scheduling concepts and multi-tenant fairness.
- Capacity planning, utilization versus saturation, fragmentation, resource overcommit and headroom.
- Heterogeneous resource modeling: CPU, memory, accelerators, NICs, topology and vendor device plugins / Dynamic Resource Allocation (DRA).
- GPU/AI batch workloads: job queues, admission, quota, gang scheduling and preemption; understand when a standard Deployment is not the right abstraction.

### Frameworks to know

- **Must know:** Kubernetes, client-go/controller-runtime/Kubebuilder, CRDs, scheduler concepts, RBAC, resource quota, device plugins and DRA concepts.
- **Study and compare:** Kueue for job admission/quotas; KEDA for event-driven autoscaling; Karpenter or the cloud's native node autoscaler; Cluster API for cluster lifecycle.
- **Use only if required:** Volcano, Ray, custom scheduler plugins, Crossplane or cloud-specific orchestration. Learn their problem boundaries instead of adopting all of them.

### Weekly sequence

| Week | Focus | Deliverable |
|---|---|---|
| 21 | Kubernetes architecture, API server/etcd/watch mechanics | Diagram a Pod from API request through scheduling to kubelet execution |
| 22 | Controller-runtime, reconciliation, caches and work queues | Operator scaffold with status conditions, rate-limited retries and metrics |
| 23 | CRD schema/versioning, finalizers, owner references and deletion | Safe resource lifecycle, cleanup, upgrade and deletion tests |
| 24 | Leader election, idempotency, optimistic concurrency and controller testing | `envtest` suite plus controller crash/restart and stale-read tests |
| 25 | Requests/limits, quotas, admission, RBAC and tenant isolation | Namespace-per-tenant policy and quota enforcement |
| 26 | Scheduler flow, filtering/scoring, affinity and topology | Simulated scheduler scoring policy and a benchmark harness |
| 27 | Queues, fairness, priorities, backfill, preemption and gang scheduling | Workload queue/admission policy with starvation and fairness tests |
| 28 | Autoscaling, capacity, fragmentation, node provisioning and disruption | Load-driven scaling experiment with capacity and tail-latency metrics |
| 29 | Device plugins, DRA and heterogeneous resource abstractions | Model CPU/GPU pools; prototype device-aware resource claims where supported |
| 30 | Multi-cluster/fleet concepts, upgrades and blast-radius control | Design a fleet rollout with staged waves, drift detection and rollback |
| 31 | **Flagship integration** | Connect the operator, queue, scheduling policy, quotas, status API and telemetry |
| 32 | **Buffer and review** | Demo failure recovery, multi-tenant fairness, scheduler results and design trade-offs |

### Phase 3 project: Workload Control Plane (flagship project begins)

Create one progressive project called **Heterogeneous Workload Control Plane**. The first version does not require real GPUs.

Core functions:
- API for submitting a workload with tenant, priority, deadline, requested resources and desired state.
- Durable lifecycle and reconciliation: Pending → Admitted → Scheduled → Running → Succeeded/Failed/Cancelled.
- Admission control, quotas, priority and a resource scheduler over simulated pools.
- Retry, timeout, worker/cluster loss recovery, status conditions and audit events.
- Metrics for queue wait, scheduling latency, utilization, fairness, failures and end-to-end completion.
- Kubernetes operator to reconcile the platform and run a real workload after simulation works.

**Exit criteria**
- Safe repeated reconciliation; no stuck finalizers on normal failure paths.
- Clear API compatibility and state-transition documentation.
- Demonstrate quota enforcement and at least two scheduler policy comparisons.
- Recover from controller restart and resource disappearance without corrupting desired state.

### Resources

- [Kubernetes architecture](https://kubernetes.io/docs/concepts/architecture/), [controller concepts](https://kubernetes.io/docs/concepts/architecture/controller/), [scheduling framework](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/)
- [Kubebuilder book](https://book.kubebuilder.io/), [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime), [Kubernetes device plugins](https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/)
- [Kueue overview](https://kueue.sigs.k8s.io/docs/overview/), [Kubernetes Dynamic Resource Allocation](https://kubernetes.io/docs/concepts/scheduling-eviction/dynamic-resource-allocation/)
- [KEDA docs](https://keda.sh/docs/), [Karpenter docs](https://karpenter.sh/docs/), [Cluster API](https://cluster-api.sigs.k8s.io/)
- *Programming Kubernetes* (O'Reilly); review upstream Kubernetes design proposals when relevant

---

## Phase 4 — Production platform engineering, SRE, delivery and security basics

**Weeks 33–42**  
**Outcome:** You can operate the control plane as a production platform, not just make it work in a local demo.

### Platform architecture and developer experience

- Platform as a product: internal customers, paved/golden paths, self-service, platform APIs and adoption feedback.
- Team Topologies: platform team boundaries, cognitive load and enabling teams.
- SOLID and design principles; composition, interfaces, dependency inversion, cohesion/coupling and explicit architecture boundaries.
- Useful patterns: adapter, strategy, state, decorator, factory where justified; CQRS/event sourcing as trade-offs, not defaults.
- RFCs/ADRs, API versioning, migration and deprecation, backward compatibility and operational documentation.
- Infrastructure as code and drift: choose one primary tool (OpenTofu/Terraform-style HCL or Pulumi in Go); understand the other at a conceptual level.
- GitOps with Argo CD or Flux; Helm/Kustomize; progressive delivery, canaries, rollback and policy gates.
- CI/CD, reproducible/hermetic builds, SBOM generation, dependency update policy, build caching and safe release practices.
- Backstage is useful for internal portals, but optional; do not confuse a portal with a platform.

### Observability and SRE

- Metrics, logs, traces and profiles; context propagation across HTTP/gRPC/messaging; OpenTelemetry Collector and semantic conventions.
- RED and USE methods, high-cardinality controls, sampling, telemetry cost and actionable alert design.
- SLIs/SLOs/error budgets, availability and latency objectives, multi-window burn-rate alerts and alert fatigue.
- Capacity modeling, load testing, queueing, autoscaling signals, saturation and bottleneck analysis.
- Incident response, ownership, on-call, runbooks, blameless postmortems, detection and mitigation.
- Disaster recovery, backups/restore testing, RTO/RPO, dependency mapping, regional failure and recovery rehearsals.
- DORA metrics are diagnostic trends, not individual productivity scores.

### Identity, security and supply chain foundations

- Cloud IAM, least privilege, secret rotation, workload identity, service accounts and mTLS.
- Kubernetes RBAC, network policies, admission policies, pod security, secrets handling and tenant boundaries.
- Threat modeling, secure defaults, patch/vulnerability management and dependency risk.
- Signed artifacts, SBOMs, provenance, SLSA concepts and image verification in CI/admission.
- Shared responsibility, data classification, audit trails, privacy, retention and residency.

### Weekly sequence

| Week | Focus | Deliverable |
|---|---|---|
| 33 | IaC and cloud networking/identity | Reproducible environment with least-privilege access and automated teardown |
| 34 | GitOps, deployment strategy, config/secrets | Git-driven deployment and a tested rollback path |
| 35 | CI, tests, supply-chain basics and artifact provenance | Build pipeline with dependency checks, SBOM and signed-image verification in a test environment |
| 36 | OpenTelemetry traces/metrics/logs and context propagation | End-to-end trace across API, queue, controller and worker |
| 37 | SLO/SLI, error budgets, alerts and load testing | SLO document, burn-rate alert and capacity-test report |
| 38 | Incidents, runbooks, failure containment and postmortems | A game day with timeline, impact, root contributing factors and corrective actions |
| 39 | Backup/restore, disaster recovery and RTO/RPO | Destroy and reconstruct the test environment; measure restore objectives |
| 40 | Platform product, golden paths, architecture and team interfaces | Developer workflow/design RFC for self-service workload onboarding |
| 41 | Tenant security, IAM, network policy and policy enforcement | Threat model plus automated least-privilege/policy tests |
| 42 | **Buffer and production-readiness review** | Release checklist, runbook, dashboard, load report and recovery demo |

### Phase 4 project: productionize the flagship platform

Deploy the Workload Control Plane into a Kubernetes test cluster using IaC and GitOps. Add CI, API/contract tests, OpenTelemetry, SLOs, load testing, a canary/rollback path, security policy and a recovery drill. The platform should be usable by someone who did not build it, using a documented deployment path.

**Exit criteria**
- Can provision and tear down the environment reproducibly.
- Have a trace from submission to completion, dashboards that answer specific questions, and alerts tied to SLOs.
- Measure recovery against declared RTO/RPO; no claim of disaster readiness without restore evidence.
- Explain tenant threat boundaries and demonstrate at least one blocked policy violation.

### Resources

- [OpenTelemetry docs](https://opentelemetry.io/docs/), [Prometheus docs](https://prometheus.io/docs/), [Grafana docs](https://grafana.com/docs/)
- [Google SRE books](https://sre.google/books/), *The SRE Workbook*, *Observability Engineering*
- [Argo CD docs](https://argo-cd.readthedocs.io/), [OpenTofu docs](https://opentofu.org/docs/), [Crossplane docs](https://docs.crossplane.io/) (compare, do not adopt all)
- [Backstage docs](https://backstage.io/docs/), [Team Topologies](https://teamtopologies.com/)
- [SLSA](https://slsa.dev/), [Sigstore](https://docs.sigstore.dev/), [Syft SBOM](https://github.com/anchore/syft), [Kyverno](https://kyverno.io/docs/), [OPA](https://www.openpolicyagent.org/docs/)
- *Platform Engineering* (Camille Fournier & Ian Nowland), *Building Secure and Reliable Systems* (Google)

---

## Phase 5 — AI compute, inference infrastructure and cost-aware orchestration

**Weeks 43–52**  
**Outcome:** You understand how model-serving demands shape compute platforms and can demonstrate a cost/latency trade-off with a real inference workload.

### AI and accelerator fundamentals

- CPU vs GPU/accelerator execution; throughput/latency trade-offs; memory bandwidth versus compute; HBM, PCIe and NVLink concepts.
- GPU memory constraints, host-to-device transfer, batching and utilization; understand roofline-style reasoning at a conceptual level.
- Prefill versus decode, time-to-first-token, inter-token latency, tokens/second, concurrency and queue wait.
- KV cache purpose, memory pressure, cache reuse/prefix caching, continuous batching and paged attention concepts.
- Quantization (for example FP8/INT8/INT4) and speculative decoding: know the quality/performance trade-offs before optimizing.
- Model lifecycle: loading, warmup, readiness, scale-out, draining, version rollout and model cache locality.
- Multi-model serving, routing, fallbacks, per-tenant/token budgets, admission, rate limits and backpressure.
- Distributed inference concepts: tensor/pipeline parallelism, all-reduce and interconnect requirements. Learn the architectural consequences first; deep collective optimization is a follow-up specialty.

### AI workload orchestration

- Batch jobs versus online serving; queue-based admission and fairness; gang scheduling concepts for distributed workloads.
- GPU/device requests, topology-aware placement, fragmentation, sharing/isolation, model residency and cold-start costs.
- Autoscaling from queue depth, request rate, tokens/sec, latency or saturation; know when each signal fails.
- Spot/preemptible instances, checkpoint/retry trade-offs and graceful interruption handling.
- Cost attribution per tenant/team/model; utilization, idle capacity, cost per request and cost per 1M tokens.
- AI observability: requests, prompt/input/output tokens, latency, errors, queue time, cache hit rates and model/resource attribution. Treat prompts and outputs as sensitive data; capture them only when justified and authorized.

### Frameworks: learn the boundary before the API

| Need | Start with | Learn after the basics |
|---|---|---|
| Model serving | vLLM or another well-supported inference runtime | SGLang/Triton comparison when useful |
| Kubernetes model deployment | KServe | A raw Deployment/Service comparison to understand what the serving layer adds |
| Batch admission/quotas | Kueue | Volcano or other schedulers if a target workload requires them |
| Device resources | Kubernetes device plugins and DRA concepts | Vendor GPU Operator and hardware-specific topology support |
| Event-driven scale | KEDA | Karpenter/node autoscaling interaction |
| Distributed execution | Understand native Kubernetes Jobs/Services first | Ray when the use case really needs distributed execution abstractions |
| LLM gateway | A minimal gateway or LiteLLM/Envoy AI Gateway | Provider/model routing, budget enforcement and failover tests |

Avoid adopting all these tools together. Pin versions in your lab and use current official documentation when implementing a feature; AI-serving ecosystems evolve quickly.

### Weekly sequence

| Week | Focus | Deliverable |
|---|---|---|
| 43 | AI serving metrics and latency taxonomy | Instrumented test with time-to-first-token, inter-token latency, throughput and error metrics |
| 44 | Deploy a small model locally or in a short-lived environment | Reproducible model-serving deployment with warmup/readiness and shutdown behavior |
| 45 | Batching, queueing, concurrency and KV cache concepts | Compare at least two serving configurations; record quality constraints and performance results |
| 46 | Model lifecycle, multi-model routing and cold starts | Model registry/version routing prototype with safe rollout and fallback |
| 47 | Kubernetes resource requests, DRA/device plugins and topology | Extend the control plane with accelerator classes and simulated topology-aware placement |
| 48 | Queue admission, quotas, fairness and batch jobs | Integrate Kueue or implement a thin experiment illustrating admission/quotas; compare policies |
| 49 | Autoscaling and scale-to-zero | Experiment with queue-based or request-based scaling; measure cold start and SLO impact |
| 50 | Cost, utilization, spot interruption and unit economics | Dashboard showing cost per request/token and handling of interrupted capacity |
| 51 | AI observability, gateway policy and tenant budgets | Traces/metrics across request → gateway → inference; budgets and rate limits tested |
| 52 | **AI infrastructure capstone review** | Architecture, performance/cost report, failure drill and prioritized backlog |

### Phase 5 project: AI workload control plane

Extend the Workload Control Plane so it can manage an inference workload. Start with a simulated accelerator pool, then connect a small real serving runtime. It should:

- Accept workload class, resource requirements, tenant, priority, latency objective and cost budget.
- Perform admission/queueing, select a resource pool and manage lifecycle.
- Surface queue wait, utilization, request latency, throughput, failures and cost estimates.
- Handle overload, failed nodes, model startup delay and interrupted capacity.
- Compare two scheduling/scaling policies under the same workload and publish methodology/results.

**Exit criteria**
- Report performance with enough context for another engineer to reproduce it.
- Explain the difference between accelerator utilization and user-facing service quality.
- Show at least one measurable trade-off among cost, queue delay, throughput and latency.
- Avoid claims of production GPU scheduling expertise based solely on a simulator; describe exactly what was measured on real hardware.

### Resources

- [vLLM documentation](https://docs.vllm.ai/), [KServe documentation](https://kserve.github.io/website/), [Kueue documentation](https://kueue.sigs.k8s.io/docs/)
- [Kubernetes scheduling framework](https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework/), [Dynamic Resource Allocation](https://kubernetes.io/docs/concepts/scheduling-eviction/dynamic-resource-allocation/)
- [NVIDIA CUDA programming guides](https://docs.nvidia.com/cuda/), [NCCL documentation](https://docs.nvidia.com/deeplearning/nccl/user-guide/docs/)
- [OpenTelemetry](https://opentelemetry.io/docs/) and its [GenAI semantic conventions repository](https://github.com/open-telemetry/semantic-conventions-genai)
- *AI Engineering* (Chip Huyen); *Designing Machine Learning Systems* (Chip Huyen) for workload and model-system context
- Read inference papers selectively: *Attention Is All You Need* for background as needed, *PagedAttention/vLLM* for serving memory management, plus current runtime design docs

---

## Phase 6 — Kernel observability, isolation, security and economics

**Weeks 53–56**  
**Outcome:** You can inspect and harden the host/runtime boundary, understand the security supply chain, and reason about infrastructure economics. This phase is intentionally focused; deeper kernel and security work is in follow-up.

### Concepts

**Linux containers and runtime isolation**
- Namespaces, cgroups v2, capabilities, seccomp, filesystem isolation and container runtime boundaries.
- CPU/memory throttling, OOM, PID limits, file descriptor limits and noisy-neighbor behavior.
- eBPF at a conceptual level: kprobes/tracepoints, TC/XDP, LSM, CO-RE, verifier constraints, observability versus enforcement.
- CNI/network policies; Cilium/Hubble as an example of eBPF-enabled networking and visibility.
- Stronger isolation options such as gVisor/Kata and their performance/compatibility trade-offs.

**Security and trust**
- Zero-trust principles, workload identity, mTLS, short-lived credentials, least privilege and secret rotation.
- Image signing, SBOM, provenance, vulnerability scanning, policy as code and admission enforcement.
- Threat modeling, tenant isolation, auditability, data minimization and incident containment.
- AI-specific security at the platform boundary: sensitive prompts/outputs, model endpoint abuse, tool permissions, prompt-injection risk in agent workflows and guarded platform actions.

**Economics and capacity**
- Cost allocation and unit economics, commitment versus on-demand/spot, interruption costs, idle resources, utilization, reservation and capacity headroom.
- Cost per request/token, cost of queueing/latency, energy-per-workload and carbon-aware scheduling as an optimization, not the first objective.
- Compliance awareness: know when obligations such as the EU AI Act, EU Cyber Resilience Act, SOC 2 or FedRAMP may apply; involve security/legal experts rather than treating a roadmap as legal advice.

### Weekly sequence

| Week | Focus | Deliverable |
|---|---|---|
| 53 | cgroups, namespaces, seccomp and container troubleshooting | Lab showing CPU/memory limits and a blocked system call in a disposable environment |
| 54 | eBPF tooling and network visibility | Trace a syscall/network event or build a read-only observation lab; do not begin with a production enforcement hook |
| 55 | Workload identity, signing, SBOM and policy-as-code | A test pipeline that rejects an unsigned image or disallowed workload |
| 56 | Cost, capacity, energy and final review | Unit-economics report plus a threat model and a 2030 skills-gap review |

### Resources

- [Linux kernel documentation](https://docs.kernel.org/), [eBPF documentation](https://ebpf.io/), [Cilium docs](https://docs.cilium.io/)
- [Kubernetes security](https://kubernetes.io/docs/concepts/security/), [SPIFFE/SPIRE](https://spiffe.io/), [gVisor](https://gvisor.dev/docs/), [Kata Containers](https://katacontainers.io/)
- [Sigstore](https://docs.sigstore.dev/), [SLSA](https://slsa.dev/), [OWASP Top 10 for LLM Applications](https://owasp.org/www-project-top-10-for-large-language-model-applications/)
- [OpenCost](https://www.opencost.io/docs/), [FinOps Foundation](https://www.finops.org/framework/)
- *Learning eBPF* (Liz Rice), *Building Secure and Reliable Systems* (Google), *Cloud FinOps* (J.R. Storment and Mike Fuller)

---

## Phase 7 — Consolidation, interviews and portfolio polish

**Weeks 57–60**  
**Outcome:** Turn the learning into credible evidence for target roles. This phase is not a substitute for continuous interview preparation; it is a final consolidation window.

| Week | Focus | Deliverable |
|---|---|---|
| 57 | Review weak areas from skills matrix and project issues | Two targeted remediation labs, chosen from actual gaps |
| 58 | HLD: workload control plane, multi-region service, inference platform | Three timed designs, each with APIs, scaling estimates, data model, failures and trade-offs |
| 59 | Coding/LLD mocks, project deep-dive and peer review | Timed Go problems, two LLD exercises, one mock project defense |
| 60 | Portfolio and roadmap review | Clean README, architecture diagrams, runbook, ADRs, benchmark report, demo and next-quarter plan |

### Core exit standard

You should be able to:
- Build and explain a reliable Go control plane and Kubernetes operator.
- Reason about consistency, failures, retries, leases, idempotency, queues and backpressure.
- Explain Kubernetes scheduling and the trade-offs among admission, placement, autoscaling and resource fairness.
- Design observability/SLOs, test restore procedures and demonstrate safe rollout/recovery.
- Run and measure an inference workload; explain latency, throughput, utilization and cost trade-offs.
- Review AI-generated code with explicit invariants, tests, profiling, security checks and failure scenarios.

Do not claim hyperscale production experience based on a local project. Describe the scale you tested, the assumptions you simulated and what would have to change at fleet scale.

---

# PART II — CONTINUOUS TRACKS (RUN ALONGSIDE THE CORE)

## A. Data structures and algorithms (DSA)

**Target:** 2–3 hours/week, usually two weekday sessions. Use Go to reduce language switching.

Cover arrays/hash maps, strings, two pointers, sliding window, stacks/queues, linked lists, binary search, trees/BSTs, heaps, graphs/BFS/DFS/topological sort, union-find, intervals, tries, recursion/backtracking, greedy, dynamic programming, bit manipulation and complexity analysis.

- Main path: [NeetCode roadmap / NeetCode 150](https://neetcode.io/roadmap)
- Reference: [Tech Interview Handbook](https://www.techinterviewhandbook.org/)
- Target: 80–120 curated problems by approximately week 29 and 150–180 reviewed by week 60, adjusted to target-company expectations.
- Process: if stuck for 20–25 minutes, inspect a hint/solution, then re-solve next day, after one week and after one month. Track the missed pattern, not just the final answer.

## B. Low-level design and code design

**Target:** one exercise every 2–3 weeks; use the flagship project to practice real design.

Topics: interfaces and contracts, SOLID principles (as heuristics, not laws), composition, cohesion/coupling, dependency injection where useful, state machines, error handling, concurrency safety, API versioning, testing seams, refactoring and design patterns. Practice rate limiter, bounded queue, worker pool, cache, scheduler, connection pool, retry library and job/workflow state machine.

- [Hello Interview LLD](https://www.hellointerview.com/learn/low-level-design/in-a-hurry/introduction)
- [Refactoring.Guru design patterns](https://refactoring.guru/design-patterns)
- [A Philosophy of Software Design](https://web.stanford.edu/~ouster/cgi-bin/book.php)
- [Awesome Low Level Design](https://github.com/ashishps1/awesome-low-level-design)

## C. High-level/system design

**Target:** one timed HLD/month, then weekly in active interview mode.

Use a consistent structure: requirements and SLOs → traffic/capacity estimates → API/data model → high-level components → consistency/failure semantics → bottleneck and scaling analysis → security → cost → trade-offs.

Practice designs directly related to the mission:
1. Durable job/workflow service.
2. Multi-tenant workload scheduler.
3. Kubernetes operator/control plane.
4. Distributed rate limiter and API gateway.
5. Multi-region control plane and fleet rollout system.
6. Object/blob metadata service or distributed cache.
7. Observability ingestion/telemetry platform.
8. GPU/LLM inference-serving platform.
9. Capacity management and autoscaling service.
10. Safe deployment and disaster recovery platform.

Resources: [System Design Primer](https://github.com/donnemartin/system-design-primer), [Hello Interview System Design](https://www.hellointerview.com/learn/system-design/in-a-hurry/introduction), DDIA and engineering papers/blogs from companies operating real infrastructure.

## D. AI-assisted engineering and verification

Use coding agents to accelerate exploration and repetitive implementation, but practice the skills that remain engineering responsibilities:
- Write specifications, invariants, failure models and acceptance tests before generating code.
- Require tests for cancellation, duplicate requests, timeouts, partial failures and resource exhaustion.
- Use race detection, fuzzing, property tests, static analysis, dependency/security scans and benchmarks.
- Review generated changes for unsafe assumptions, API compatibility, resource leaks, over-retries, security boundaries and observability gaps.
- Keep a record of where AI output was wrong and how you detected it.

Resource: [MIT Missing Semester](https://missing.csail.mit.edu/) and language/framework documentation. Revisit your employer's security and data-handling policies before sending proprietary code or logs to any assistant.

## E. Research reading, writing and open source

- Read one foundational paper every two weeks when in the distributed-systems/platform phase; write a half-page summary and the design lesson.
- Read one upstream design proposal or production engineering post per month.
- Write one ADR/RFC for each project milestone, plus a short postmortem after each failure drill.
- Start contributing after Phase 2: docs/tests first, then bug fixes. Aim for meaningful contributions rather than a vanity PR count.
- Potential projects: Kubernetes/controller-runtime, Kueue, OpenTelemetry, Cilium, KServe, vLLM or the tool you actually use.

### Foundational papers and designs

Read these selectively as their concepts appear in the plan:
- Lamport: *Time, Clocks, and the Ordering of Events*; FLP; *Paxos Made Simple*.
- Ongaro/Ousterhout: *In Search of an Understandable Consensus Algorithm* (Raft).
- Dynamo, GFS, Bigtable, Spanner, Borg, Omega, Chubby, Dapper and *The Tail at Scale*.
- *Kafka: a Distributed Messaging System for Log Processing*; CRDT foundations.
- *Metastable Failures in Distributed Systems*; *Use of Formal Methods at AWS*.
- *What Every Programmer Should Know About Memory* (Ulrich Drepper) as a reference, not a memorization exercise.
- *PagedAttention/vLLM* and current inference system papers when studying serving internals.

---

# PART III — FOLLOW-UP TRACKS (DEFERRED, BUT NOT FORGOTTEN)

These tracks are sequenced intentionally. Finish the control-plane core first; then select one specialty based on the role you want. Do **not** take all tracks at once. Rust is included as the first follow-up because you explicitly want it, but it should not displace the core Go/Kubernetes work.

## Follow-up 1 — Rust systems programming (recommended first, 8–12 weeks)

**Prerequisites:** Comfortable with Go concurrency, interfaces, memory/performance profiling and the core control-plane project. You may start Rust syntax earlier for curiosity, but do the serious systems track after the core project has momentum.

### Sequence

1. **Rust fundamentals (2–3 weeks):** ownership, borrowing, lifetimes, structs/enums, traits, generics, `Result`/`Option`, iterators, modules, error handling, testing and Cargo.
2. **Memory/concurrency (2 weeks):** `Send`/`Sync`, atomics, mutexes, channels, pinning/futures concepts, safe versus unsafe boundaries and how the compiler enforces aliasing rules.
3. **Async and networking (2–3 weeks):** Tokio runtime, tasks, cancellation, bounded channels, backpressure, timeouts, graceful shutdown, async I/O and structured error handling. Understand async trade-offs rather than assuming it is always faster.
4. **Performance and tooling (1–2 weeks):** Criterion benchmarks, flamegraphs/profiling, allocation behavior, `cargo clippy`, formatting, tests and fuzzing.
5. **Integration project (2 weeks):** replace one measured hot-path component of the Go platform with a Rust sidecar/library or implement a small ingestion/compaction worker using a stable gRPC/Protobuf boundary. Benchmark both implementations under the same conditions.

### Resource order

- [The Rust Programming Language](https://doc.rust-lang.org/book/) + [Rustlings](https://github.com/rust-lang/rustlings)
- [Rust by Example](https://doc.rust-lang.org/rust-by-example/)
- *Rust for Rustaceans* (Jon Gjengset)
- *Rust Atomics and Locks* (Mara Bos)
- [Asynchronous Programming in Rust](https://rust-lang.github.io/async-book/) and [Tokio tutorial](https://tokio.rs/tokio/tutorial)
- [Criterion.rs](https://bheisler.github.io/criterion.rs/book/)

**Exit criteria:** A tested Rust component with clear ownership boundaries, no unexplained `unsafe`, meaningful benchmarks and a written explanation of why Rust was justified. Do not rewrite the whole platform in Rust without evidence.

## Follow-up 2 — Deep storage engines and database internals (6–10 weeks)

Take this if targeting storage systems, databases, metadata services or infrastructure persistence.

**Sequence:**
1. WAL, durability and crash recovery.
2. Memtables/SSTables, LSM trees, compaction, tombstones, bloom filters and write/read amplification.
3. B-tree/B+ tree, page layout, buffer pool and indexes.
4. MVCC, isolation, locking, checkpointing, vacuum and query plans.
5. Replication, snapshots, backup/restore, log retention and rebalancing.
6. Benchmark write/read mixes, compaction debt, tail latency and recovery behavior.

**Project:** a small storage-engine lab with WAL, recovery and one index structure—or a benchmark/prototype that modifies an established engine. This is explicitly a learning project, not production database software.

**Resources:** DDIA, *Database Internals* (Alex Petrov), [CMU Database Group courses](https://15445.courses.cs.cmu.edu/), [PostgreSQL internals](https://www.postgresql.org/docs/), [Pebble](https://github.com/cockroachdb/pebble), [RocksDB](https://github.com/facebook/rocksdb).

## Follow-up 3 — Streaming, durable workflows and data infrastructure (4–8 weeks)

Take this when a target job operates event streams, telemetry, change-data capture, lakehouses or durable business workflows.

**Sequence:** partitioning/order → consumer groups and rebalances → offsets/replay → outbox/CDC → event-time/watermarks/state → backfill/reprocessing → durable workflows/sagas → end-to-end tracing and operational recovery.

**Project:** Postgres outbox/CDC → Kafka or Redpanda → a small stateful stream processor → an analytical table, with replay/backfill and idempotency tests. Add Temporal only if a multi-step process needs durable execution.

**Resources:** [Kafka docs](https://kafka.apache.org/documentation/), [Debezium](https://debezium.io/documentation/), [Apache Flink](https://nightlies.apache.org/flink/flink-docs-stable/), [Apache Iceberg](https://iceberg.apache.org/), [Temporal](https://docs.temporal.io/).

## Follow-up 4 — Kernel, eBPF and host security (6–10 weeks)

Take this for infrastructure performance, observability, CNI, runtime security or node-level engineering roles.

**Sequence:** Linux process/memory/I/O internals → tracing with `perf`/bpftrace → eBPF maps/programs/verifier/CO-RE → tracepoints/kprobes → TC/XDP/CNI fundamentals → LSM/security enforcement → profiling and overhead measurement → production rollout and fallback strategy.

**Project:** start with read-only latency/connection observability; only later prototype a narrowly scoped enforcement policy in an isolated VM. Never experiment with unreviewed enforcement code on a production host.

**Resources:** [Linux kernel docs](https://docs.kernel.org/), [eBPF](https://ebpf.io/), [Cilium](https://docs.cilium.io/), *Learning eBPF* (Liz Rice), *BPF Performance Tools* (Brendan Gregg).

## Follow-up 5 — Advanced networking and low-latency data planes (6–12 weeks)

Take this if aiming at cloud networking, storage data paths, AI cluster networking or performance-intensive infrastructure.

**Sequence:** TCP/congestion/control and Linux networking → routing and ECMP concepts → proxy/L7 internals (Envoy) → CNI and eBPF → NUMA and NIC queues → RDMA/RoCE concepts → collective communication/NCCL concepts → traffic shaping, congestion and end-to-end latency measurement.

BGP/FRR, MPLS, EVPN/VXLAN, DPDK, XDP and deep RDMA implementation are **role-dependent**. Know the concepts that explain connectivity and topology; specialize only if job descriptions point in this direction.

**Project:** profile and improve a network-heavy service, or build a small Envoy extension/traffic policy and measure tail latency. For AI networking, compare serving performance with different topology/communication assumptions before attempting distributed GPU benchmarks.

**Resources:** [Envoy docs](https://www.envoyproxy.io/docs), [Cilium docs](https://docs.cilium.io/), [Linux networking docs](https://docs.kernel.org/networking/), [NCCL docs](https://docs.nvidia.com/deeplearning/nccl/user-guide/docs/), [DPDK docs](https://doc.dpdk.org/guides/).

## Follow-up 6 — Advanced AI serving and distributed training (8–12 weeks)

Take this after Phase 5 when the role involves high-volume inference, model-serving performance or AI cluster scheduling.

**Sequence:** profiling basics → model-serving runtime internals → batching/KV cache → quantization/speculative decoding/prefix caching → tensor/pipeline parallelism → NCCL collectives and topology → disaggregated prefill/decode → GPU memory management → autoscaling under SLOs → workload-level cost model.

**Project:** make a reproducible before/after performance improvement in an inference serving path. Report quality constraints, TTFT, inter-token latency, tokens/sec, memory, concurrency and cost; explain hardware and traffic assumptions.

**Resources:** [vLLM](https://docs.vllm.ai/), [SGLang](https://docs.sglang.ai/), [NVIDIA CUDA](https://docs.nvidia.com/cuda/), [NCCL](https://docs.nvidia.com/deeplearning/nccl/user-guide/docs/), [KServe](https://kserve.github.io/website/), current serving papers and upstream design docs.

## Follow-up 7 — Advanced formal methods and distributed verification (4–8 weeks)

Take this for correctness-critical systems, databases, consensus, schedulers or control planes with complex concurrency.

**Sequence:** write explicit invariants → TLA+ state/action models → fairness and liveness assumptions → model checking → history checking/linearizability → deterministic simulation → property-based and fuzz testing → CI regression corpus.

**Project:** model and test a lease/failover or scheduler admission protocol with failures and concurrency; compare the model's assumptions to implementation behavior.

**Resources:** [TLA+](https://lamport.azurewebsites.net/tla/tla.html), [Apalache](https://apalache-mc.org/), [Porcupine](https://github.com/anishathalye/porcupine), [Jepsen](https://jepsen.io/analyses), [FoundationDB papers/talks](https://www.foundationdb.org/).

## Follow-up 8 — Platform agents, MCP and AI operations (2–4 weeks; optional)

Treat this as a product interface on top of a well-designed platform, not as a substitute for a control plane.

**Sequence:** tool APIs and typed schemas → identity and permissions → allowlisted actions → dry-run/approval for destructive changes → audit logs and rate limits → prompt-injection/threat testing → evaluation and rollback.

**Project:** expose read-only status/SLO queries first, then a tightly scoped deploy/rollback action with authorization, approval and audit history. The underlying controller remains responsible for correctness.

**Resources:** [Model Context Protocol](https://modelcontextprotocol.io/), [OWASP GenAI/LLM security](https://owasp.org/www-project-top-10-for-large-language-model-applications/), [OpenTelemetry GenAI semantic conventions](https://github.com/open-telemetry/semantic-conventions-genai).

## Follow-up 9 — Staff-level architecture and technical leadership (ongoing)

After the core and one specialty, practice technical direction: problem framing, RFCs, cross-team APIs, migration plans, multi-quarter roadmaps, incident leadership, cost/reliability trade-offs and mentoring. Read *Staff Engineer* (Will Larson), *The Staff Engineer's Path* (Tanya Reilly) and *Team Topologies*. The goal is not just to build components, but to make systems and teams easier to operate.

---

# PART IV — CONCEPT COVERAGE CHECKLIST

Use this as a completeness audit. A checkmark means the topic is covered in the roadmap; it does not imply expert-level mastery in every item.

## Systems and performance

- [ ] CPU/memory hierarchy, caching, NUMA concepts, allocation, virtual memory and page cache
- [ ] Processes/threads, scheduling, syscalls, locks, atomics, memory ordering and concurrency safety
- [ ] Go scheduler, garbage collection, context cancellation, channels, race detector, profiling
- [ ] Linux networking/I/O, `epoll`, `io_uring` concepts, zero-copy trade-offs
- [ ] Benchmark design, percentiles, tail latency, flamegraphs, CPU/heap/contention profiles

## Networking and interfaces

- [ ] DNS, TCP lifecycle, retransmission/congestion, TLS, HTTP/1.1–3, QUIC, L4/L7 load balancing
- [ ] RPC/gRPC, Protobuf evolution, deadlines, retries, idempotency and streaming
- [ ] API lifecycle/versioning, compatibility, pagination, errors and rate limiting
- [ ] Proxy/gateway, CNI, CSI and network policy concepts
- [ ] Follow-up: Envoy internals, eBPF networking, BGP/ECMP, RDMA/RoCE, DPDK when relevant

## Distributed systems and data

- [ ] Partial failure, safety/liveness, time/ordering, CAP/PACELC and FLP assumptions
- [ ] Linearizability, eventual/causal consistency, transactions, MVCC and isolation
- [ ] Consensus/Raft, leader election, leases/fencing, quorum, replication and membership
- [ ] Partitioning, sharding, rebalancing, multi-region, replication lag and data residency
- [ ] WAL, B-tree/LSM, compaction, indexes, snapshots and recovery concepts
- [ ] Messaging delivery semantics, idempotency, outbox/CDC, replay, event time and workflows
- [ ] Retries, jitter, budgets, circuit breakers, bulkheads, load shedding, queueing and caching
- [ ] Property tests, fuzzing, TLA+, deterministic simulation, chaos and linearizability checking

## Control plane and platform

- [ ] API server/etcd/controller/scheduler/kubelet/CRI/CNI/CSI architecture
- [ ] Reconciliation, caches/watches, resource versions, finalizers, status conditions and owner refs
- [ ] Operator testing, leader election, idempotency, upgrades and recovery
- [ ] Admission, RBAC, tenancy, quotas, resource requests/limits and policy
- [ ] Scheduling filters/scoring, placement, topology, fairness, priorities, preemption and gang concepts
- [ ] Autoscaling, capacity, fragmentation, DRA/device plugins and cluster/fleet lifecycle
- [ ] IaC, GitOps, progressive delivery, CI, release safety, golden paths and platform product thinking
- [ ] Observability, SLOs, incident management, disaster recovery, RTO/RPO and capacity planning

## AI infrastructure

- [ ] CPU/GPU/accelerator trade-offs, HBM, memory bandwidth, PCIe/NVLink and topology concepts
- [ ] Prefill/decode, TTFT/inter-token latency, batching, KV cache, throughput and model lifecycle
- [ ] Admission/queues, quotas, GPU resource management, topology-aware placement and fairness
- [ ] Serving runtime, autoscaling, model routing, gateways, tenant/token budgets and fallbacks
- [ ] Quantization/speculative decoding/prefix caching and distributed-inference concepts
- [ ] Cost per request/token, spot interruption, utilization, capacity headroom and energy trade-offs

## Security, operations and professional engineering

- [ ] Cloud IAM, least privilege, workload identity, mTLS, secrets, isolation and threat modeling
- [ ] cgroups/namespaces/seccomp, supply chain, signatures, SBOM, provenance and policy-as-code
- [ ] AI data privacy, prompt/tool risks, guarded action interfaces and auditability
- [ ] FinOps unit economics; compliance awareness appropriate to target markets
- [ ] SOLID/design principles, useful design patterns, ADR/RFC writing, DSA, LLD and HLD
- [ ] AI-assisted coding with explicit specs, test verification, code review and performance/security checks
- [ ] Open-source contribution, technical writing, incident postmortems and stakeholder communication

---

# PART V — CAREER EXECUTION AND MEASUREMENT

## Portfolio checklist

By the end of the core plan, aim for:

- [ ] One progressive **Heterogeneous Workload Control Plane** repository
- [ ] One reliable Go dispatcher and benchmark report
- [ ] Durable state and failure semantics documented and tested
- [ ] Kubernetes operator with lifecycle, quotas, scheduling and recovery
- [ ] End-to-end traces, dashboards, SLOs, runbooks and a measured restore drill
- [ ] AI inference integration with a reproducible cost/performance comparison
- [ ] Threat model, security checks and resource teardown/budget controls
- [ ] 4–6 meaningful ADRs/RFCs, 2–4 concise technical write-ups and 2 postmortems
- [ ] At least one useful upstream contribution after you understand the project
- [ ] DSA/LLD/HLD review log and several mock interviews

A smaller, rigorous project is better than many incomplete repositories. Public artifacts should not contain employer code, secrets, user data or unapproved cloud details.

## Role-selection filter

When evaluating a job, look for evidence that the team owns one or more difficult systems:

- Distributed APIs/control planes and lifecycle management
- Resource scheduling, queues, quotas, placement, fleet management or capacity
- Kubernetes internals, operators, multi-cluster or cloud infrastructure
- Performance, reliability, failure recovery or high-volume data paths
- AI inference/compute infrastructure with measurable latency, utilization and cost objectives

Ask interviewers what proportion of the work is platform internals versus integration/configuration; how they measure success; what failures the team owns; how often engineers design and review RFCs; and whether the platform runs at meaningful scale. Do not choose a company based on the words “AI” or “platform” alone.

## Quarterly plan review

Every 12–13 weeks:
1. Review 10–15 real job descriptions for your target roles.
2. Identify repeated missing concepts, not one-off tool names.
3. Review the official docs/release notes for the frameworks you actually use.
4. Reassess project scope and drop low-value extras before adding new technologies.
5. Record concrete evidence of improvement: test coverage of invariants, recovery time, p99 latency, throughput, utilization, cost per request and design quality.

## Manager-facing alignment (if useful)

Frame the plan as business outcomes rather than personal study:

| Capability | Potential outcome | Evidence/metric |
|---|---|---|
| Profiling/performance | More throughput per compute unit | CPU/memory per request, p99 latency |
| Control-plane correctness | Fewer stuck or duplicated operations | Reconciliation success, retry/failure rate, manual intervention count |
| Scheduling/capacity | Better utilization and fairer service | Queue delay, fragmentation, utilization, starvation and SLO attainment |
| Observability/SRE | Faster diagnosis and recovery | MTTD/MTTR, SLO attainment, alert quality |
| Platform automation | Safer and faster delivery | Lead time, change failure rate, rollback time |
| AI infrastructure | Better inference economics | Cost per request/1M tokens, TTFT, throughput and accelerator utilization |
| Security | Reduced operational exposure | Policy violations prevented, signed artifact coverage and vulnerability remediation time |

Ask for a modest, time-boxed study allocation, a safe sandbox with spending limits and opportunities to present demonstrable outcomes. Do not assume the employer will accept the entire plan; use it to negotiate a small initial project aligned to team goals.

---

## Final guidance

This roadmap aims for **breadth of conceptual coverage and depth in the target specialty**. It deliberately sequences the work so that control-plane engineering, distributed-systems correctness, scheduling and production operations arrive before advanced language/runtime or low-level specialties. Rust is a planned follow-up, not an abandoned idea. Storage engines, full streaming pipelines, deep eBPF, advanced networking, distributed GPU optimization and agentic operations remain available in dependency order for later specialization.

No roadmap can guarantee that every framework or technology important in 2030 is known today. What makes this plan future-resilient is the recurring practice of mapping new tools onto stable ideas: state, invariants, scheduling, failure, isolation, performance, observability, economics and operability.
