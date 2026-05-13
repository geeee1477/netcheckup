# netcheckup

Lightweight network diagnostic CLI written in Go.

`netcheckup` helps analyze common connectivity problems by checking:

- DNS resolution
- ICMP/Ping reachability
- TCP connectivity
- HTTP/HTTPS availability

The tool provides both human-readable summaries and machine-readable JSON output for automation and scripting.

---

# Features

- DNS resolution checks
- Ping/ICMP testing
- TCP port connectivity testing
- HTTP/HTTPS health checks
- Concurrent execution for faster diagnostics
- Retry support
- Configurable timeouts
- Timing metrics for all checks
- JSON output mode
- Quiet mode
- Exit codes for automation and CI/CD
- IPv4 preference for primary IP detection

---

# Installation

## Run locally

```bash
go run cmd/netcheckup/main.go google.com
```

## Install globally

```bash
go install github.com/geeee1477/netcheckup/cmd/netcheckup@latest
```

After installation:

```bash
netcheckup google.com
```

---

# Usage

```bash
netcheckup [options] <target>
```

---

# Examples

## Basic check

```bash
netcheckup google.com
```

## Custom port

```bash
netcheckup --port 80 google.com
```

## Custom timeout

```bash
netcheckup --timeout 2 google.com
```

## Retry failed checks

```bash
netcheckup --retries 3 google.com
```

## JSON output

```bash
netcheckup --json google.com
```

## Quiet mode

```bash
netcheckup --quiet google.com
```

---

# Options

| Flag | Description | Default |
|---|---|---|
| `--port` | TCP/HTTP port to test | `443` |
| `--timeout` | Timeout in seconds | `3` |
| `--retries` | Number of retries | `2` |
| `--json` | Output JSON instead of text | `false` |
| `--quiet` | Hide verbose logs | `false` |

---

# Example Output

## Human-readable output

```text
========== SUMMARY ==========
✔ DNS resolution works (28 ms)
✔ Host reachable via ping (27 ms)
✔ TCP connection successful (36 ms)
✔ HTTP service responding (199 ms)

→ Target is fully reachable and functioning
```

---

## JSON output

```json
{
  "target": "google.com",
  "primary_ip": "74.125.29.139",
  "dns_ok": true,
  "ping_ok": true,
  "tcp_ok": true,
  "http_ok": true,
  "dns_ms": 28,
  "ping_ms": 27,
  "tcp_ms": 36,
  "http_ms": 199,
  "diagnosis": "target_fully_reachable"
}
```

---

# Exit Codes

| Exit Code | Meaning |
|---|---|
| `0` | All checks successful |
| `1` | One or more checks failed |

---

# Project Structure

```text
netcheckup/
├── cmd/netcheckup/
│   └── main.go
├── internal/checks/
│   ├── dns.go
│   ├── ping.go
│   ├── tcp.go
│   ├── http.go
│   └── summary.go
└── README.md
```

---

# Planned Features

- Traceroute support
- Packet loss analysis
- Multi-target scanning
- TLS certificate inspection
- DNS server selection
- Port scan mode
- Export formats
- Monitoring integrations

---

# License

MIT License
