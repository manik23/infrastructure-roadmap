# Prerequisite Cheatsheet

**Roadmap source:** [Entry assessment and lab environment](../../ROADMAP_2030_REVISED.md#3-entry-assessment-and-lab-environment)

## Linux triage order

1. Listener and owner: `ss -ltnp` or `lsof -nP -iTCP:<port> -sTCP:LISTEN`.
2. Process resources: `lsof -p <pid>` and `/proc/<pid>/fd`.
3. Behavior: trace a safe disposable process with `strace`; distinguish blocking calls from errno failures.
4. Logs: `journalctl -u <unit>` for a systemd service; `docker logs <container>` for a container without systemd.

## Request path

`DNS → TCP handshake → TLS handshake → encrypted HTTP request/response`.

- Immediate TCP refusal commonly means no listener or an explicit reject.
- A timeout can occur in DNS, connect, TLS, or response wait.
- `401`/`403` and `404` are HTTP responses, not timeouts.

## Memory and safety

- Page cache is reclaimable cached file data in RAM; it is not a page table.
- RSS is resident memory attributed to a process; page cache is system-wide file-backed memory.
- Before cloud/GPU work: owner, budget alert, quota, expiry/teardown, and a CPU/simulation fallback.
