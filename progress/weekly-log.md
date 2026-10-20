# Weekly / Session Log

Append a new entry for each meaningful study session. Prefer one entry per session, then summarize into a weekly review. Do not overwrite prior evidence.

---

## Template
- **Date:** YYYY-MM-DD
- **Phase / week:**
- **Objective:**
- **Resource and section:**
- **Concepts / invariants:**
- **Work completed:**
- **Evidence:** (tests, output, benchmark, design artifact; say `not run` when appropriate)
- **Commands run:**
- **AI assistance used:** (what was generated/suggested and how it was verified)
- **What I can explain unaided:**
- **Gaps / incorrect assumptions found:**
- **Next action:**
- **Git checkpoint:** (branch, commit, synced/pushed status; verify at session start)

---

## Initial checkpoint
- **Date:** Not started
- **Phase / week:** Prerequisite assessment
- **Objective:** Establish evidence-based placement
- **Work completed:** None
- **Evidence:** None
- **Next action:** Follow `prompts/00-START-PREREQUISITE-ASSESSMENT.md`

---

## Session checkpoint
- **Date:** 2026-10-19
- **Phase / week:** Prerequisite assessment → targeted prerequisite block
- **Objective:** Establish an evidence-based placement before Phase 1 and verify the safe local lab environment.
- **Resource and section:** `ROADMAP_2030_REVISED.md` entry assessment; roadmap-coach prerequisite-assessment protocol.
- **Concepts / invariants:** Linux service triage; DNS → TCP → TLS → HTTP ordering; file-backed page cache versus RSS; Go cancellation/shutdown; idempotent Kubernetes finalizer cleanup; least-privilege IAM; priority queues; Git bisection; container runtime diagnosis.
- **Work completed:** Interactive diagnostic responses across all eight baseline domains; assessment, targeted prerequisite plan, lab-readiness checklist, and placement records updated. Local Docker, Go, and Git availability verified.
- **Evidence:** Docker Engine `linux/arm64 29.8.2`; Alpine container reported Linux `7.0.14-linuxkit` on `aarch64`; Go `1.27.1` on darwin/arm64; Git `2.54.0`. Baseline topic statuses and observed gaps are recorded in `progress/BASELINE_ASSESSMENT.md`.
- **Commands run:** `docker version --format '{{.Server.Os}}/{{.Server.Arch}} {{.Server.Version}}'`; `docker run --rm alpine:3.20 uname -a`; `go version`; `git --version`; `git pull --rebase --autostash` (fast-forwarded `main` to `1c11214`, then local progress-file conflict resolved); `git diff --cached --check` (passed).
- **AI assistance used:** Diagnostic questions, explanations, and a proposed disposable Docker Linux-triage lab. Suggestions were not treated as evidence without learner response or command output.
- **What I can explain unaided:** Partial evidence only: hash-set duplicate detection, max-heap selection, basic Go cancellation concerns, and Kubernetes deletion-state awareness.
- **Gaps / incorrect assumptions found:** Linux process/socket/syscall triage; HTTPS failure classification; page-cache/RSS distinction; Go worker lifecycle; Kubernetes `404` cleanup idempotency; cloud IAM/cost controls; Git bisection; container runtime/CI smoke tests.
- **Next action:** Run the proposed disposable Linux HTTP-service triage lab and capture listener, FD, `/proc`, request, and container-log evidence. The lab itself is **not run**.
- **Git checkpoint:** `main` aligned with `origin/main` at `1c11214` before local progress edits. Five progress files are staged locally; no commit or push was made. `stash@{0}` remains as a recoverable autostash safety copy.

---

## Session checkpoint
- **Date:** 2026-10-20
- **Phase / week:** Targeted prerequisite block — Linux service triage
- **Objective:** Add a roadmap-linked revision system and obtain the first local HTTP-service observation.
- **Resource and section:** `ROADMAP_2030_REVISED.md` entry assessment/lab environment; `revision/ROADMAP_MAP.md`.
- **Concepts / invariants:** Revision documents are retrieval aids, not curriculum authority. A successful HTTP response proves request delivery and application response, not process/socket/FD/syscall diagnosis.
- **Work completed:** Added the `revision/` map, notes, and cheatsheets; linked the folder from the README. Performed one request against the disposable local HTTP server.
- **Evidence:** `curl -i http://127.0.0.1:18080/` returned `HTTP/1.0 200 OK` from `SimpleHTTP/0.6 Python/3.12.13`. Listener/PID, FDs, `/proc`, `strace`, and container-log evidence: **not run**.
- **Commands run:** Learner: `curl -i http://127.0.0.1:18080/`. Agent: `git fetch --prune` (succeeded); `git diff --check` (passed); `graphify ROADMAP_2030_REVISED.md --no-viz` (failed: CLI expects a directory); `graphify . --no-viz` (stopped: installed version requires an external semantic-extraction key for 15 docs).
- **AI assistance used:** Created a native Markdown map after the optional graph tool could not complete without an external key. The map was checked with `git diff --check`; no semantic graph was produced.
- **What I can explain unaided:** Not reassessed this session.
- **Gaps / incorrect assumptions found:** Linux triage remains incomplete; a 200 response alone is insufficient evidence for listener, FD, syscall, or logging diagnosis.
- **Next action:** Run `docker exec roadmap-linux-triage ss -ltnp` and interpret the listener/PID output before collecting FD evidence.
- **Git checkpoint:** `baseline_tested` matched `origin/baseline_tested` at `0828d14` before this session's local edits. No commit or push was made.

---

## Session checkpoint
- **Date:** 2026-10-20
- **Phase / week:** Targeted prerequisite block — Linux triage, HTTPS request path, OS memory/I/O
- **Objective:** Produce bounded, observable evidence for service triage, HTTPS phase recognition, and page-cache/reclaim reasoning without privileged, destructive, cloud, or GPU work.
- **Resource and section:** `ROADMAP_2030_REVISED.md` entry assessment and prerequisite resources. External reading: **not needed** for these short exercises. Revision notes were updated with the completed scope; Go implementation did not begin.
- **Concepts / invariants:** A listener FD may be healthily idle in a readiness wait; a request accepted into a connected FD should progress through read/log/write or return an errno. HTTPS follows DNS → TCP → TLS certificate verification → encrypted HTTP. RSS is process-resident memory; page cache is kernel-managed cached file data; clean cache may be reclaimed, while dirty cache needs writeback to backing storage first.
- **Work completed:** Completed the disposable container Linux HTTP-service observation: listener/PID, `lsof`, `/proc` FDs, Docker logs, syscall trace, `ppoll` readiness trace, and published-port mapping. Captured two bounded verbose HTTPS client requests. Created/read/removed a 64 MiB temporary cache-lab file. Added detailed notes and concise cheatsheets. Added `scripts/check-revision-roadmap-alignment.sh` and wired it into pre-commit to guard the derived Phase 1–7 map/notes against heading/order drift.
- **Evidence:**
  - Linux **Practiced** — `python3` PID 1 owned listener FD 3 on `*:8080`; `/proc/1/fd` correlated socket inode `10373`; request trace read 78 bytes, logged 61 bytes, sent 156-byte headers and an 805-byte body, and shut down without displayed errno. Idle `ppoll` timeouts and a `POLLIN` readiness event were observed. `docker port` confirmed `8080/tcp -> 127.0.0.1:18080`.
  - Networking **Demonstrated** (baseline happy path) — `curl -v` resolved `example.com` to `104.20.23.154` and `172.66.147.243`, connected to `:443`, completed TLS 1.3 with certificate verification, negotiated HTTP/2 via ALPN, sent `GET /`, and received `HTTP/2 200`; body-printing request received 577 bytes.
  - OS, memory and I/O **Practiced** — explained RSS/page-cache/reclaim correctly after correction. The cache lab showed `MemAvailable: 7452240 kB`, `Cached: 2363908 kB`, and `SReclaimable: 270752 kB`; both reads rounded to `0.00s`, so no caching-performance claim is made. The 64 MiB file was removed.
  - Kernel OOM logs **Not assessed** — ordinary container `dmesg` returned no usable lines with exit status 0; Docker Debug returned `dmesg: read kernel buffer failed: Operation not permitted`. No OOM was induced and no privileged Debug shell was used.
  - Go concurrency **Practiced** (prior code-review evidence only); Kubernetes controllers, cloud/IAM/cost, Git/containers/CI remain **Review needed**; DSA remains **Practiced**; Rust remains **Deferred by design**.
- **Commands run:**
  - Learner: `curl -i http://127.0.0.1:18080/`; `docker exec roadmap-linux-triage ss -ltnp`; `docker exec roadmap-linux-triage lsof -nP -iTCP:8080 -sTCP:LISTEN`; `docker exec roadmap-linux-triage sh -c 'ls -l /proc/1/fd'`; `docker logs roadmap-linux-triage`; `docker exec -it roadmap-linux-triage sh -c 'strace -f -tt -T -e trace=network,read,write -p 1'`; `docker debug roadmap-linux-triage`; `strace -f -tt -T -y -p 1 -e trace=network,read,write,poll,ppoll,select,pselect6,epoll_wait,epoll_pwait,accept,accept4,connect`; `docker port roadmap-linux-triage`; `curl -v --connect-timeout 5 --max-time 15 https://example.com/ -o /dev/null`; `curl -v --connect-timeout 5 --max-time 15 https://example.com/`; `docker exec roadmap-linux-triage sh -c 'dd if=/dev/zero of=/tmp/page-cache-lab bs=1M count=64 conv=fsync && grep -E "^(MemAvailable|Cached|SReclaimable):" /proc/meminfo && time sh -c "cat /tmp/page-cache-lab >/dev/null" && time sh -c "cat /tmp/page-cache-lab >/dev/null"'`; `docker exec roadmap-linux-triage rm -f /tmp/page-cache-lab`; `docker exec roadmap-linux-triage sh -c 'dmesg | tail -n 50'`; `echo $?`; and `dmesg | tail -n 30` inside Docker Debug.
  - Agent: `docker debug --help`; Graphify query attempted but no `graphify-out/graph.json` existed; documented-heading inspection; `./scripts/check-revision-roadmap-alignment.sh` (**passed**); `git diff --check` (**passed**); `git status --short --branch` (**recorded below**).
- **Tests / benchmarks:** Automated Go tests, race tests, Kubernetes tests, CI smoke tests, and benchmarks: **not run**. The two timed `cat` reads are an observation only, not a benchmark.
- **Graphify:** **skipped / not run**. No semantic backend was explicitly approved and available for this session; no graph artifacts were refreshed, reviewed, or staged.
- **AI assistance used:** The agent supplied bounded diagnostic commands, explained observed traces, proposed the cache lab, and drafted progress/revision records. Learner-provided command output was used as evidence; no generated implementation or test suite was accepted without verification.
- **What I can explain unaided:** High-level HTTPS lifecycle and common boundaries; RSS versus page cache; clean versus dirty cache reclaim; Linux listener/FD/syscall/log triage order. Precision for idle listener versus a blocked handler still needs a later fresh scenario.
- **Gaps / incorrect assumptions found:** `ppoll` was initially described as a watchdog and later an idle listener as blocked; `ping` was initially proposed as TCP/HTTP verification. HTTPS initially skipped TCP between DNS and TLS. Page cache was initially conflated with page tables, and dirty writeback was initially described as possibly going to virtual memory. Corrections are captured in the revision notes.
- **Next action:** Define bounded Go worker-pool cancellation, input-closure, result/error ownership, and worker-completion semantics; then implement the targeted lab and run `go test -race`.
- **Git checkpoint:** Branch `baseline_tested`; latest local commit before this checkpoint is `f35419b`. Working tree contains local, uncommitted roadmap/progress/revision changes and one new alignment-check script. No commit or push was performed. **LOCAL / NOT YET SHARED.**
