# Prerequisite Revision Notes

**Roadmap source:** [Entry assessment and lab environment](../../ROADMAP_2030_REVISED.md#3-entry-assessment-and-lab-environment)
**Status:** In progress — prerequisite block needed; see [baseline assessment](../../progress/BASELINE_ASSESSMENT.md).

## Current evidence

- Docker Desktop provides a disposable Linux/arm64 environment; Go and Git are verified locally.
- Linux triage, HTTPS failure classification, page-cache/RSS reasoning, cloud IAM/cost safety, Git bisection, and container smoke testing need focused practice.
- Go concurrency and basic DSA have partial evidence only.
- Rust is deferred by design; no early-Rust prerequisite was found.

## Evidence to add before changing status

- Command output and interpretation for the Linux HTTP-service triage lab.
- A labeled HTTPS request trace and safe failure-classification exercise.
- A bounded worker-pool test including cancellation and `-race` evidence.
- A finalizer cleanup test, IAM policy-simulation/design result, and Git/container workflow artifacts.

## Linux triage exercise — listener ownership

**Problem:** A local HTTP service responds, but an incident responder must identify which process owns the listening socket before investigating FDs, syscalls, or logs.

**Core model:** A listening TCP socket is a file descriptor owned by a process. `ss -ltnp` reports the listening endpoint and, where permissions permit, the owning process and FD.

**Command breakdown:** `ss` is the Linux socket-statistics utility (the modern successor to much `netstat` usage). `-l` selects listening sockets, `-t` limits output to TCP, `-n` keeps numeric addresses/ports instead of resolving names, and `-p` requests the owning process/PID and file-descriptor information. Process details may require sufficient permission for processes owned by other users.

**Why it matters:** Control-plane and service incidents often begin with “is the intended process actually listening?” Identifying the listener/PID establishes the boundary for `/proc`, `lsof`, `strace`, and logs.

**Observed command and signal (2026-10-20):**

```text
docker exec roadmap-linux-triage ss -ltnp
LISTEN 0 5 0.0.0.0:8080 0.0.0.0:* users:(("python3",pid=1,fd=3))
```

**Interpretation:** `python3` (PID 1) owns listening socket FD 3 on port 8080 inside the container. `0.0.0.0:8080` means all container interfaces; Docker's `127.0.0.1:18080` host mapping is a separate boundary. `Recv-Q 0` shows no pending accepted connections at the observation instant. The reported backlog is 5.

**Limits and next evidence:** This does not show per-process non-socket FDs, blocking syscalls, request handling, or logs. Capture `lsof`, `/proc/<pid>/fd`, a request, and `docker logs` before reassessing Linux triage.

## Linux triage exercise — listening FD verification

**Problem:** Verify that the PID found through socket inspection truly owns the expected listening file descriptor.

**Core model:** Unix processes access sockets through file descriptors. `lsof` maps a process, FD number, socket protocol, address, and state into one record.

**Why it matters:** Corroborating the listener with a second tool avoids diagnosing the wrong process and connects network symptoms to the process-level FD table.

**Observed command and signal (2026-10-20):**

```text
docker exec roadmap-linux-triage lsof -nP -iTCP:8080 -sTCP:LISTEN
python3  1  root  3u  IPv4  10373  0t0  TCP *:8080 (LISTEN)
```

**Interpretation:** PID 1 (`python3`) runs as `root` inside this disposable container and owns FD 3. `3u` means the descriptor is open for read/write. It is an IPv4 TCP socket in `LISTEN` state on every IPv4 container interface (`*:8080`), matching the preceding `ss` output. `SIZE/OFF 0t0` is not a response-size measurement; sockets do not have ordinary file length semantics.

**Limits and next evidence:** This command selects only the listening TCP FD. Inspect `/proc/1/fd` to see the full descriptor table, then observe a request and container logs; no syscall behavior has been measured yet.

## Linux triage exercise — full descriptor table

**Problem:** Distinguish the listening socket from the process's standard streams and identify where container logs originate.

**Core model:** `/proc/<pid>/fd` exposes a process's descriptor table as symbolic links. Socket links use an inode-like identifier that can be correlated with socket-inspection tools; stdout/stderr often flow through container-runtime pipes.

**Why it matters:** FD exhaustion, closed standard streams, wrong log routing, and leaked sockets are process-level causes of service failures. The table is a quick boundary check before deeper tracing.

**Observed command and signal (2026-10-20):**

```text
docker exec roadmap-linux-triage sh -c 'ls -l /proc/1/fd'
0 -> /dev/null
1 -> pipe:[14561]
2 -> pipe:[14562]
3 -> socket:[10373]
```

**Interpretation:** The service has no interactive stdin (`FD 0` is `/dev/null`). FDs 1 and 2 are write-only pipes, which Docker captures as container stdout/stderr logs. FD 3 is socket inode `10373`, matching the socket identifier from `lsof`; this correlates the Python process's descriptor table with its listening TCP socket.

**Limits and next evidence:** A four-FD snapshot does not establish that the process will not leak descriptors under load or that it handles requests correctly. Inspect `docker logs` to connect the earlier HTTP request to output, then use a safe syscall observation to distinguish request I/O from blocking/failure behavior.

## Linux triage exercise — container log correlation

**Problem:** Confirm that a client request reached the application and determine the correct log boundary for a container without systemd.

**Core model:** Container runtimes capture the main process's stdout/stderr streams. `docker logs <container>` retrieves those captured streams; it is the appropriate log view for this simple container instead of `journalctl`.

**Command breakdown:** `docker logs` asks Docker for a container's captured standard output/error. `roadmap-linux-triage` is the container name; no flags were needed for the complete current log. Use `-f` only when following new output continuously.

**Why it matters:** A listening socket alone does not prove an application handled a request. Correlating the request with an application log distinguishes network reachability from application-level handling.

**Observed command and signal (2026-10-20):**

```text
docker logs roadmap-linux-triage
172.17.0.1 - - [20/Oct/2026 04:28:42] "GET / HTTP/1.1" 200 -
```

**Interpretation:** The Docker bridge-side client address (`172.17.0.1`) sent `GET /` over HTTP/1.1; Python handled it and returned status `200`. The log is emitted through the stdout/stderr pipes previously seen as FDs 1 and 2.

**Limits and next evidence:** This is one successful request, not a latency, error-rate, or syscall analysis. Observe system calls around another request to distinguish socket accept/read/write behavior from a process blocked in I/O.

## Linux triage exercise — syscall request path

**Problem:** Determine whether the listening service accepts a request, reads it, emits its log, sends a response, and closes the connection without a syscall-level failure.

**Core model:** A TCP server accepts a connection from its listening socket into a new connected FD, reads request bytes, writes logs and response bytes, then closes or half-closes the connected FD. `strace` observes the kernel boundary; its per-syscall duration is not end-to-end request latency.

**Command breakdown:** `strace` traces system calls; `-f` follows child processes/threads; `-tt` adds timestamps; `-T` records syscall duration; `-e trace=network,read,write` filters to request-I/O operations; `-p 1` attaches to the Python server. Attaching may be blocked by container ptrace policy; do not add privileges merely to force this lab.

**Why it matters:** A service can listen yet block, fail reads/writes, or emit errors. The trace establishes where the request crossed from network to application behavior and whether a kernel-visible error occurred.

**Observed command and signal (2026-10-20):**

```text
accept4(3, ... 172.17.0.1:58248 ...) = 4
getsockname(4, ... 172.17.0.2:8080 ...) = 0
recvfrom(4, "GET / HTTP/1.1...", 8192, ...) = 78
write(2, "172.17.0.1 - - ...", 61) = 61
sendto(4, "HTTP/1.0 200 OK...", 156, ...) = 156
sendto(4, "<!DOCTYPE HTML>...", 805, ...) = 805
shutdown(4, SHUT_WR) = 0
```

**Interpretation:** Listening FD 3 accepted a Docker-bridge client into connected FD 4. The service read 78 request bytes, wrote 61 bytes to FD 2 (the Docker-captured log pipe), wrote a 156-byte response header and the 805-byte HTML body, then cleanly closed its write side. The body length matches the earlier `Content-Length: 805`. All displayed operations returned non-negative values, so this trace contains no observed errno failure. The separate traced execution context is expected for the threaded HTTP server.

**Limits and next evidence:** This is one happy-path request under negligible load. It does not establish behavior under timeouts, broken clients, FD exhaustion, overload, or sustained latency. Before changing Linux status beyond `Practiced`, explain this trace unaided and propose the safe diagnostic order for a slow service.

## Linux triage exercise — idle wait versus blocked I/O

**Problem:** Decide whether a service that appears inactive is stuck on I/O or is correctly waiting for the next connection.

**Core model:** Event-driven and server loops use readiness waits such as `poll`/`ppoll` to sleep until a file descriptor is ready or a timeout expires. A wait on the listening socket is healthy when no clients are pending. A long wait becomes evidence worth investigating only when work is expected to be ready but cannot progress.

**Command breakdown:** `docker debug roadmap-linux-triage` opens Docker's diagnostic toolbox shell without changing the application image. The trace command used `-f` to follow the handler child, `-tt` for timestamps, `-T` for syscall duration, `-y` to label FDs, `-p 1` to attach to the server, and an explicit syscall filter for socket I/O and readiness waits. Only one tracer can attach to a process at a time; the earlier tracer had to be stopped with `Ctrl-C` first. This disposable-container trace is read-only; end it with `Ctrl-C`, then `exit` the debug shell.

**Why it matters:** In production triage, an idle readiness wait must not be mistaken for a hung service. The useful question is whether a request makes the FD ready and then progresses through accept, read, and response writes, or instead stays blocked or returns an errno.

**Observed command and signal (2026-10-20):**

```text
strace -f -tt -T -y -p 1 \
  -e trace=network,read,write,poll,ppoll,select,pselect6,epoll_wait,epoll_pwait,accept,accept4,connect

ppoll([{fd=3<socket:[10373]>, events=POLLIN}], 1, {tv_sec=0, tv_nsec=500000000}, ...) = 0 (Timeout) <0.50...>
ppoll([{fd=3<socket:[10373]>, events=POLLIN}], 1, {tv_sec=0, tv_nsec=500000000}, ...) = 1 ([{fd=3, revents=POLLIN}], ...) <0.287901>
accept4(3<socket:[10373]>, ...) = 4<socket:[8757]> <0.000324>
[pid 108] recvfrom(4<socket:[8757]>, ...) = 78 <0.000164>
[pid 108] sendto(4<socket:[8757]>, ...) = 156 <0.000112>
[pid 108] sendto(4<socket:[8757]>, ...) = 805 <0.000093>
```

**Interpretation:** FD 3 is the earlier listener socket. The repeated approximately 0.5-second `ppoll` timeouts are healthy idle waiting: no connection was queued. At `05:01:16`, `POLLIN` means the listener became readable because a connection arrived; it returned before the 0.5-second deadline. PID 1 accepted it as FD 4, handler PID 108 immediately read the 78-byte request and sent the same response observed earlier. The interleaved `unfinished`/`resumed` lines show concurrent parent and handler syscalls, not a fault. No displayed call has a negative return or unexpectedly long duration, so this is not blocked-I/O evidence.

**Limits and next evidence:** This proves only that this low-load request progressed. A slow service can still block on a different FD, CPU, lock, dependency, or resource limit. For a real slow-request investigation, first correlate a pending request, then inspect listener/PID, FD table, relevant wait/error syscalls, and logs in that order.

**Common diagnostic trap:** ICMP `ping` checks whether an address answers ICMP under its network policy; it does not prove that the intended TCP port, Docker port mapping, HTTP route, or application listener is reachable. When `ppoll` on a listener continues to time out while a TCP client waits, verify the exact TCP endpoint and path with the client result, listener/port mapping, and service logs before attributing it to a firewall.

**Interpretation rules from the learner's assessment:** `ss` reporting no listener is one local cause of a client-side connection refusal, but only if traffic reaches that endpoint; a network drop commonly appears as a timeout instead. `/proc/<pid>/fd` confirms descriptor ownership, while `lsof -p <pid>` and resource-limit/error evidence are needed to investigate leakage or exhaustion. A readiness timeout is meaningful only when timestamps prove a request is pending. Correlate the client timestamp with both `strace` and the application boundary (`docker logs` here, `journalctl` for a systemd service).

**Next safe path check:** When the client uses a Docker-published host port but the container listener never becomes ready, run `docker port roadmap-linux-triage`. `docker port` displays the configured host-to-container published-port mapping and does not alter the container. Confirm it maps the client target (for this lab, host `127.0.0.1:18080`) to container port `8080`, then retry the exact `curl` target and correlate its timestamp. A missing/wrong mapping is a more direct hypothesis than a firewall; investigate host/daemon network policy only after the endpoint and mapping agree.

**Observed mapping (2026-10-20):** `docker port roadmap-linux-triage` returned `8080/tcp -> 127.0.0.1:18080`, matching the earlier successful host-side `curl`. This confirms the configured local forwarding for the lab; it is not by itself evidence about a different client network.

## Networking assessment — HTTPS phases and failure boundaries

**Problem:** Locate a client-visible request failure in the correct protocol phase instead of treating every failure as an application timeout.

**Core model:** For a typical HTTPS URL, the client resolves the hostname with DNS, establishes a TCP connection to the chosen IP and port, performs the TLS handshake and certificate verification, then sends encrypted HTTP and receives an HTTP response. Each phase has distinct evidence and owners.

**Why it matters:** Infrastructure engineers must avoid debugging an HTTP handler when the TCP connection never formed, or calling a received `401` a timeout. Accurate phase classification narrows the responsible boundary: resolver, network/load balancer, TLS configuration, or application.

**Assessment signals (2026-10-20):** The learner identified common causes: no listener as one cause of refusal, invalid/expired certificates as TLS failure, authentication as the meaning of `401`, and a slow/stuck handler as one cause of a response timeout. The complete phase ordering and boundaries remain under assessment.

**Interpretation rules:** A connection refusal normally means the endpoint actively rejected the TCP connection after it was reached (commonly no listener, but an explicit reject is also possible); a silent packet drop instead tends toward a connect timeout. Certificate verification happens during TLS, after TCP but before HTTP; expiry, hostname mismatch, an untrusted issuer/chain, or invalid time can cause it. An HTTP `401` is a received application-layer response, not a transport failure. A response timeout must be placed precisely: it can be waiting for TLS, for the request write, or—only after an HTTP request was sent—for headers/body because of server, proxy, downstream, or network delay.

**Next evidence:** Explain the phases unaided, then capture a verbose local client trace and label where its observable events fall. No external reading is needed before this short assessment; use *High Performance Browser Networking* DNS, TCP, TLS, and HTTP sections only if this remains unclear after the exercise.

**Observed safe trace (2026-10-20):** The learner ran `curl -v --connect-timeout 5 --max-time 15 https://example.com/ -o /dev/null`, then a second `curl -v --connect-timeout 5 --max-time 15 https://example.com/` that printed the body. `curl` is the HTTP client; `-v` prints connection, TLS, request, and response metadata (not decrypted application payload); `--connect-timeout 5` limits only connection establishment to five seconds; `--max-time 15` bounds the whole transfer; `-o /dev/null` discards the response body while preserving the terminal trace. Both are read-only public HTTPS requests with no cloud/GPU cost. The trace reported hostname resolution to `104.20.23.154` and `172.66.147.243`, connection to `104.20.23.154:443`, TLS 1.3, a verified `example.com` certificate, ALPN selection of `h2`, `GET / HTTP/2`, and `HTTP/2 200`; the body-printing request received 577 bytes. HTTP/2 formatting and any proxy behavior are client/environment dependent; this does not prove HTTP/3 or packet-level TCP behavior.

## OS-memory assessment — RSS, page cache, and reclaim

**Problem:** Explain why a process can have modest resident memory while the machine holds gigabytes of recently read file data in RAM, and predict the behavior when RAM becomes scarce.

**Core model:** RSS is resident physical memory mapped into a particular process, including anonymous and file-backed pages currently resident for that process. The page cache is kernel-managed RAM holding cached file contents, potentially shared by many readers and not owned exclusively by one process. It is cached data, not a page table or a count of virtual addresses.

**Why it matters:** Large control-plane logs, images, checkpoints, and model artifacts can make memory look full even when application RSS is moderate. Incorrectly blaming a process for shared cache leads to bad capacity decisions and unnecessary restarts.

**Reclaim model:** Under pressure, Linux generally reclaims clean, unused cache pages because they can be re-read from storage. Dirty file-cache pages must be written back before they can be reclaimed. Accessing a file page after it was evicted creates a page fault; if storage must supply it, that is a major fault and can add I/O latency. Pressure does not itself mean every process immediately faults.

**Assessment evidence (2026-10-20):** The learner correctly separated approximately 100 MiB of process RSS from several GiB of kernel page cache, and then correctly stated that clean cache is reclaimable, dirty cache requires writeback, and access after eviction can fault. “Disk/virtual memory” was corrected to writeback to the backing file on storage.

**Next safe observation:** In the disposable container, run `docker exec roadmap-linux-triage sh -c 'dd if=/dev/zero of=/tmp/page-cache-lab bs=1M count=64 conv=fsync && grep -E "^(MemAvailable|Cached|SReclaimable):" /proc/meminfo && time sh -c "cat /tmp/page-cache-lab >/dev/null" && time sh -c "cat /tmp/page-cache-lab >/dev/null"'`. `dd` creates a 64 MiB temporary file from zero bytes; `bs=1M` and `count=64` bound its size, and `conv=fsync` waits for that write to reach the file's storage layer. `grep` reads selected kernel memory counters; the two `cat` commands read the same file, discarding output, and shell `time` reports wall-clock time. Expect the second read to be served from cache when the container/host has sufficient memory, but do not claim a speedup from noisy timing alone. Docker Desktop's VM and shared host cache limit attribution: this does not expose macOS host memory or prove a specific process's RSS. Cleanup is explicit: `docker exec roadmap-linux-triage rm -f /tmp/page-cache-lab` after recording the output. Do not induce an OOM event. Use focused OSTEP virtualization and persistence material only if the corrected model does not hold; no external reading is required before this bounded lab.

**Observed output (2026-10-20):** `dd` wrote 67,108,864 bytes in 0.027759 seconds and reported 2.3 GB/s. `/proc/meminfo` reported `MemAvailable: 7452240 kB`, `Cached: 2363908 kB`, and `SReclaimable: 270752 kB`. Both timed reads rounded to `0.00s`. This is consistent with a warm cached read but does not measure a first-versus-second-read difference at useful resolution; the reported write rate is not physical-disk throughput evidence because the container uses Docker Desktop's VM and overlay filesystem. `docker exec roadmap-linux-triage rm -f /tmp/page-cache-lab` then completed with no output, which is expected and confirms temporary-file cleanup.

**Planned read-only OOM check:** Run `docker exec roadmap-linux-triage sh -c 'dmesg | tail -n 30'`. `dmesg` reads the kernel message buffer; `tail -n 30` limits output to the newest 30 lines. A container commonly lacks permission to read the host kernel buffer, and an accessible buffer may contain no OOM event; record either outcome as evidence. Do not create memory pressure or kill processes to manufacture an OOM result.

**OOM-check outcome (2026-10-20):** `docker exec roadmap-linux-triage sh -c 'dmesg | tail -n 50'` printed no kernel lines and `echo $?` returned `0`. Docker Debug then reported `dmesg: read kernel buffer failed: Operation not permitted`. Taken together, these commands do not supply reliable kernel OOM evidence; mark that subtopic `Not assessed` in Docker Desktop rather than inferring no OOM occurred. Do not add `--privileged` solely for this assessment. Revisit `journalctl -k` or permitted `dmesg` in a dedicated Linux VM after a naturally occurring event.

## Go concurrency assessment — bounded worker-pool ownership

**Problem:** A caller cancels a bounded job-dispatch operation while workers may be receiving jobs, running work, or trying to report results. The design must not leak goroutines or deadlock if the caller stops reading results.

**Core model:** Channel closure has one owner: the sender side that knows no more values will be produced. A worker receiving from a job channel must also select on `ctx.Done()`; a worker sending a result must be able to stop that send on cancellation. A coordinating goroutine waits for all workers before it closes the results channel. Cancellation does not retroactively undo a job that has already been handed to a handler; the API must report whether that handler observed the context and how unfinished queued work is classified.

**Why it matters:** Controller and scheduler workers routinely stop during leader changes, shutdown, retry cancellation, and deadline expiry. Ownership and cancellation rules define whether the system exits cleanly or accumulates blocked goroutines under load.

**Planned review exercise:** Inspect a deliberately incomplete Go worker-pool sketch in the study session. Identify which producer, workers, or coordinator can remain blocked if the caller cancels and stops consuming results. No external reading is required before this first review attempt; before implementation, read the Go `context` package overview plus `WithCancel` and `WithTimeout` documentation. Expected later evidence is a runnable bounded implementation and `go test -race` result; neither has been run yet.

**Assessment evidence (2026-10-20):** The learner correctly saw that the sketch did not select cancellation before invoking `handle`, but initially concluded that all jobs would run. In fact, if the caller cancels and stops receiving, a worker can block at `results <-`; blocked workers stop receiving jobs, the bounded producer can block at `jobs <-`, the job channel is never closed, and `wg.Wait()` cannot finish. This is a liveness failure. The next review must cover cancellation at both receive and result-send boundaries.

**Channel ownership evidence (2026-10-20):** The learner correctly assigned `close(jobs)` to the routine that sends jobs and `close(results)` after all workers finish. The latter must be a coordinator after `wg.Wait()`, not a worker: a different worker may still send and would panic if the shared results channel were closed early.

**Cancellation-policy decision (2026-10-21):** Choose Option A for the first implementation: cancellation stops admission and abandons queued, unstarted jobs to avoid unnecessary compute. The producer selects on `ctx.Done()` while sending, then closes `jobs`. Each dequeued handler receives the same context; it can stop early only if it cooperates. Because workers also select on cancellation while sending results, cancellation may prevent delivery of an in-flight outcome. The API must return `ctx.Err()` and document that absent results are incomplete/abandoned, not successful. This is an at-least-once-attempt boundary, not a guarantee that every submitted job has an outcome record.

**Required reading before implementation:** [Go `context` package](https://pkg.go.dev/context), bounded to its overview and the `WithCancel` and `WithTimeout` function documentation. Purpose: distinguish cancellation propagation from a forceful stop, and make the worker-pool deadline boundary explicit.

**Context review evidence (2026-10-21):** The learner correctly stated that `cancel()` closes `Done` and makes `Err()` report cancellation; it does not stop a goroutine. The per-job timeout context belongs in the worker immediately before `handle(jobCtx, job)`, derived from the pool context with `context.WithTimeout`. The worker calls the returned `jobCancel()` immediately after `handle` returns—do not `defer` it across a long-lived worker loop—so its timer/resources are released even on a normal completion. A handler must observe the context and choose to return `ctx.Err()`; cancellation itself does not fabricate a `Result` or error return.

**Job-error policy decision (2026-10-21):** An ordinary error returned by `handle` is emitted in that job's result and does not cancel the pool. This keeps independent job failure isolated; only parent cancellation/deadline ends the pool. The implementation must distinguish this per-job error from `ctx.Err()`, which means the overall operation was canceled and queued jobs may have been abandoned.

**Lab scaffold (2026-10-21):** Created [labs/prerequisites/workerpool](../../labs/prerequisites/workerpool) in the root Go module as a bounded prerequisite lab, with requirements, assumptions, failure model, and acceptance tests in its README. It is explicitly not the Phase 1 dispatcher; no worker-pool implementation or behavior test exists yet.

**Scaffold verification (2026-10-21):** The initial go test ./... check passed
for the root module's one package with no test files. It verifies module/package
wiring only; worker-pool behavior and race tests are not run.

**Specific exercise and review gate (2026-10-21):** The worker-pool README now
defines the target API and deterministic acceptance tests, beginning with one
ordinary job error followed by successful later work. REVIEW.md defines
test-first increments, channel/context review checks, exact verification
commands, and separate safety/liveness review. No behavior implementation or
acceptance test has been added yet.

**Learner-owned implementation scaffold (2026-11-02):**
[`pool.go`](../../labs/prerequisites/workerpool/pool.go) contains only the
public contract (`Job`, `Result`, `Handler`, `Config`, and a panic-TODO `Run`).
The learner implements the first ordinary-job-error increment in that bounded
file; they do not need to type unit-test boilerplate. The problem is to ensure
that one ordinary handler error is represented only in that job's result and
does not stop later independent work. This distinction matters because workers
in a dispatcher or controller must distinguish per-item failure from
operation-wide cancellation. After the learner submits the implementation, the
agent adds a deterministic test using one worker and jobs `a`, `b`, `c`, with a
sentinel error only for `b`, plus appropriate edge cases. A green result will
prove only ordinary error isolation, not bounded concurrency, cancellation,
deadline propagation, or race freedom. Status: **Review needed** until the
implementation, generated tests, and their output are reviewed.

**AI-driven unit-test policy (2026-11-02):** Unit tests and fixtures are
generated by the agent after each learner-written implementation increment. The
learner's task is to understand and write the core logic, then explain its
invariant and review the generated tests and edge cases. Generated tests remain
proposals: a passing result is evidence only after the learner can explain what
the assertions prove and what they do not prove.

**Card 1 implementation and review evidence (2026-11-02):** The learner
implemented sequential `Run` in
[`pool.go`](../../labs/prerequisites/workerpool/pool.go): for each input job it
calls `handle(ctx, job)` and appends the resulting ID, value, and handler error
to a `Result`; it returns `nil` as the operation error. The reviewer added
[`pool_test.go`](../../labs/prerequisites/workerpool/pool_test.go): an
ordinary-error test confirms that `a`, `b`, `c` are handled in order and that
only `b` carries a sentinel error; a context-value test permits the later
deadline child context while requiring caller values to propagate; an empty
input test prevents accidental handler invocation. `gofmt` and `git diff
--check` completed without output. `GOCACHE=/private/tmp/roadmap-workerpool-go-build-cache go test ./labs/prerequisites/workerpool` returned `ok` in `0.587s`.
This is **Practiced**, not demonstrated: the tests do not establish worker
count, queue bounds, cancellation, deadline behavior, race safety, retries,
persistence, or panic containment.

**Card 1 explanation evidence (2026-11-02):** The learner correctly explained
that returning `boom` as the `Run` operation error would make this API
fail-fast: job `c` and later queued jobs would not run, and their results would
be absent. Under this lab's policy, that is wrong because ordinary job failure
is a per-item outcome, while `Run` errors represent pool-level failure such as
invalid configuration or later parent cancellation. Card 1 is therefore
**Demonstrated**; the broader worker-pool topic remains **Practiced** pending
Card 2 concurrency and later cancellation/deadline/race evidence.

**Card 2 learner scaffold (2026-11-02):** The TODO block directly inside
[`Run`](../../labs/prerequisites/workerpool/pool.go) is the scaffold for
evolving the verified Card 1 sequential baseline in place. It exposes the five
decision points without providing their logic: configuration validation; queue
allocation; worker/result production; producer close ownership; and completion
coordination. The expected behavior and acceptance criteria are in Card 2 of
`REVIEW.md`; no Card 2 behavior evidence exists yet.

**Card 2 review finding (2026-11-02):** The learner added a bounded jobs
channel, `cfg.Workers` worker goroutines, producer-owned `close(jobQueue)`,
configuration validation, and `wg.Wait()`. The reviewer added tests for an
active-handler bound, all-job completion, and invalid configuration. The code
then had every worker execute `results = append(results, ...)` on the same
slice. A slice append can mutate its shared header and backing array, so this
requires a mutex or, preferably here, a single result collector receiving from
a worker-to-collector channel. `go test` and two observed `go test -race` runs
passed, but that does not establish safety: a dynamic race detector only sees
executed schedules, while the absence of synchronization is a correctness
defect by inspection. Card 2 status is **Review needed**. The fix must state
one result-owner and preserve liveness: producer closes jobs, workers exit,
then the coordinator safely closes/finalizes the result path after `wg.Wait()`.

**Card 2 result-ownership revision (2026-11-02):** The learner replaced the
shared slice append with a `results` channel buffered to the caller-input
length. Each worker sends one `Result`; after `wg.Wait()`, the coordinator
closes that channel and a single collection loop appends to the returned slice.
`go test` passed in `1.180s`, `go test -race` passed in `1.637s`, and `git diff
--check` passed. This resolves the shared-slice ownership defect for the
observed Card 2 path. However, the revision also derived a
`context.WithTimeout` child context and called its cancel function. That is
Card 4 deadline behavior and has no corresponding propagation test yet; it
must be removed for Card 2 rather than accepted early. Status remains **Review
needed**.

**Card 2 scope-correction evidence (2026-11-02):** The learner removed the
premature per-job `context.WithTimeout` behavior, leaving Card 2 to pass the
parent context directly. Fresh uncached `go test -count=1` passed in `0.728s`
and `go test -race -count=1` passed in `1.637s`; `gofmt -d` printed no changes
and `git diff --check` passed. The result channel is buffered to the finite
input length, so workers can send one result each while `Run` waits. The
learner must still explain why that buffer avoids a Card 2 blocking cycle and
why closing before `wg.Wait()` would race with a worker send. Status:
**Practiced**, pending explanation for **Demonstrated**.

**Card 2 explanation evidence (2026-11-02):** The learner correctly explained
that the result channel has one slot per finite input job, so each worker can
publish one result while `Run` waits for worker completion. `wg.Wait()` proves
no worker remains that can send before the coordinator closes results; closing
it earlier could panic with `send on closed channel`. Card 2 bounded concurrency
and result ownership are **Demonstrated**. The full worker-pool topic remains
**Practiced** because cancellation, deadline propagation, and repetition/race
hardening are still open.

**Card 3 scaffold (2026-11-02):** The in-place TODO block in
[`Run`](../../labs/prerequisites/workerpool/pool.go) identifies the cancellation
boundaries: job admission, job receive, and result send. Its invariant is that
parent cancellation stops new admission and returns partial results with the
parent terminal error; missing results mean unstarted or undelivered work, not
success. A handler already running receives the parent context and must choose
to cooperate. The planned test uses channels to ensure a handler has started
before cancellation; it must not use `time.Sleep`. Card 3 status: **Not
assessed** until learner implementation and review tests exist.

**Card 3 partial-review evidence (2026-11-02):** The learner made producer
admission select on `ctx.Done()` and close `jobQueue` before returning on
cancellation. Agent-authored tests now cover cancellation after a handler
starts and an already-canceled context with two workers; fresh normal and race
runs passed in `0.231s` and `1.906s`. The learner then made the worker receive
directly cancellation-selectable. Card 3 remains **Review needed** because its
result send is still plain `results <- result` after a `ctx.Err()` check. A
buffered result channel happens to avoid blocking in this finite-slice
topology, but the stated invariant is stronger: every potentially blocking
operation must be in a `select` with `ctx.Done()`. Checking `ctx.Err()` before
or after an operation leaves a time-of-check/time-of-block gap; cancellation
can arrive after that check while the goroutine blocks.

**Card 3 receive-selection edge case (2026-11-02):** The learner subsequently
made result send cancellation-selectable and changed worker receive to a
`select`. Fresh normal tests passed (`0.980s`), but the fresh race run failed
the deterministic cancellation test because a queued handler began after
cancellation. This is not a data race report: it is `select` semantics. If
both `ctx.Done()` and a buffered job receive are ready, Go may choose either
case. The worker must therefore check cancellation again after receiving a job
and before invoking `handle`, so a queued job cannot begin after the
abandon-queued-work policy has taken effect. Card 3 remains **Review needed**.

**Card 3 post-receive guard evidence (2026-11-02):** The learner added a
post-receive `ctx.Err()` guard before `handle`, so a worker that selected a
queued job at the same time cancellation became ready does not start it. The
producer admission, worker receive, and result send now each have a `select`
with `ctx.Done()`. Fresh `go test -count=1` passed in `0.232s`, and fresh
`go test -race -count=1` passed in `1.636s`. Card 3 is **Practiced** pending
the learner's explanation of shutdown and partial-result semantics.

**Card 3 explanation evidence (2026-11-02):** The learner correctly explained
that the producer's cancellation case stops admission and closes `jobQueue`;
workers terminate either by selecting cancellation before receive, by the
post-receive cancellation guard, or when the closed jobs channel is drained;
and `wg.Wait()` ensures no worker can send before results close. The learner
also explained that a result may be absent not because it is invalid, but
because cancellation intentionally chooses partial-result delivery after an
in-flight handler outcome. Card 3 is **Demonstrated**.

**Card 4 scaffold (2026-11-02):** The in-place TODO block in
[`Run`](../../labs/prerequisites/workerpool/pool.go) separates parent-pool
cancellation from a child handler deadline. The child is derived immediately
before `handle` with `context.WithTimeout`; its cancel function must run as
soon as that handler returns, avoiding a loop-wide deferred timer leak. A child
`DeadlineExceeded` is a per-job result error when the parent remains active;
it is not a `Run` error. Card 4 status: **Not assessed** until implementation
and generated deadline test evidence exist.

**Card 4 implementation evidence (2026-11-02):** The learner derived a child
`context.WithTimeout` immediately before `handle` and called its cancel
function immediately afterward. The reviewer added a test where the handler
waits on its child context; fresh normal and race runs passed in `0.741s` and
`1.946s`. The test observed a handler deadline and `context.DeadlineExceeded`
in its job result while `Run` returned a nil operation error because the parent
context remained active. Card 4 is **Practiced** pending explanation of the
two error scopes and immediate timer cleanup.

**Card 4 explanation evidence (2026-11-02):** The learner correctly explained
that a child context limits only one job attempt, so its deadline becomes that
job's `Result.Err` rather than canceling the parent `Run` operation. The learner
also explained that calling the child cancel function immediately releases its
timer/resources, whereas a defer in a long-running worker loop would retain
them until the worker exits. Card 4 is **Demonstrated**.

**Card 5 hardening task (2026-11-02):** Run `go test -race
./labs/prerequisites/workerpool -count=50 -timeout=10s`. A passing run raises
confidence for the exercised tests but cannot prove universal race freedom or
stop a handler that ignores its context. Record that latter limitation before
declaring the worker-pool prerequisite complete.

**Card 5 verification evidence (2026-11-02):** `go test -race
./labs/prerequisites/workerpool -count=50 -timeout=10s` passed in `2.823s`.
This is repeated dynamic evidence for the current tests, not a proof that every
possible interleaving is race-free. It also cannot stop a handler that ignores
its context; the worker remains in `handle` until that function returns. Card
5 is **Practiced** pending the learner's explanation of this limitation.

**Card 5 limitation explanation (2026-11-02):** The learner correctly
explained that a timer changes the supervising goroutine's wait/policy, but
does not forcibly terminate a blocked handler. A context-ignoring handler keeps
its worker inside `handle` and can delay pool shutdown. Cards 1–5 are therefore
**Demonstrated** for this targeted worker-pool prerequisite, with the explicit
handler-cooperation limitation retained for future dispatcher/control-plane
design.

## Learner notes

Add only concise, unaided explanations here. Link the relevant `progress/` evidence; do not copy the full roadmap or mark an exit test complete without proof.
