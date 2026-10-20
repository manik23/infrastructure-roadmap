# Prerequisite Cheatsheet

**Roadmap source:** [Entry assessment and lab environment](../../ROADMAP_2030_REVISED.md#3-entry-assessment-and-lab-environment)

## Linux triage order

1. Listener and owner: `ss -ltnp` or `lsof -nP -iTCP:<port> -sTCP:LISTEN`.
2. Process resources: `lsof -p <pid>` and `/proc/<pid>/fd`.
3. Behavior: trace a safe disposable process with `strace`; distinguish blocking calls from errno failures.
4. Logs: `journalctl -u <unit>` for a systemd service; `docker logs <container>` for a container without systemd.

`ss -ltnp`: `ss` = socket statistics; `-l` = listening; `-t` = TCP; `-n` = numeric addresses/ports; `-p` = owning process/PID/FD when permitted.

### Readiness-wait interpretation

- `ppoll([{fd=<listener>, events=POLLIN}], ...) = 0 (Timeout)`: no connection became ready before the deadline; normal when idle.
- `... = 1 ([{fd=<listener>, revents=POLLIN}], ...)`: a connection is waiting; expect `accept4` next.
- A handler blocked after `accept4` is different from an idle listener with no `accept4`.
- A pending client plus listener timeouts means the request did not reach that listener. Check the exact client target and, for Docker, `docker port <container>` before blaming the firewall.
- Only one `strace` instance can attach to a process. End the previous trace with `Ctrl-C`; then use `Ctrl-C` and `exit` to leave a read-only Docker Debug session.

## Request path

`DNS → TCP handshake → TLS handshake → encrypted HTTP request/response`.

- Immediate TCP refusal commonly means no listener or an explicit reject.
- A timeout can occur in DNS, connect, TLS, or response wait.
- `401`/`403` and `404` are HTTP responses, not timeouts.
- TLS verification is after TCP and before HTTP; check expiry, hostname, issuer/chain trust, and local time.

### `curl -v` phase markers

| Trace text | Meaning |
|---|---|
| `Host ... was resolved` | DNS produced an address. |
| `Trying <ip>:443` then `Connected to` | TCP connection established. |
| `TLS handshake`, `SSL connection using`, `certificate verify ok` | TLS negotiated and certificate verification succeeded. |
| `ALPN: server accepted h2` | Application protocol negotiation selected HTTP/2. |
| `> GET ...` | HTTP request sent over TLS. |
| `< HTTP/2 200` | HTTP response received; `200` is application success. |

Safe bounded trace: `curl -v --connect-timeout 5 --max-time 15 https://example.com/ -o /dev/null`.

## Memory and safety

- Page cache is reclaimable cached file data in RAM; it is not a page table.
- RSS is resident memory mapped into a process; page cache is system-wide kernel-managed file-backed memory and can be shared.
- Under pressure, clean cache can be evicted; dirty cache must be written back to its backing file on storage first. A later access to an evicted file page can trigger a major fault and storage I/O.
- Kernel OOM logs may be inaccessible from Docker Desktop containers. Record `Not assessed`; do not add `--privileged` or induce an OOM just to obtain baseline evidence.
- Before cloud/GPU work: owner, budget alert, quota, expiry/teardown, and a CPU/simulation fallback.
