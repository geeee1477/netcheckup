![Go Version](https://img.shields.io/badge/go-1.26-blue)
![License](https://img.shields.io/badge/license-MIT-green)

# netcheckup

Lightweight network diagnostic CLI written in Go.

`netcheckup` helps analyze common connectivity problems by checking:

- DNS resolution
- ICMP/Ping reachability
- TCP connectivity
- HTTP/HTTPS availability
- TLS certificate health
- Packet loss
- Network routing paths
- Open TCP ports

The tool provides both human-readable summaries and machine-readable JSON output for automation, scripting, monitoring, and CI/CD workflows.

---

# Features

- DNS resolution checks
- Custom DNS server support
- ICMP/Ping reachability testing
- TCP connectivity testing
- HTTP/HTTPS health checks
- TLS certificate inspection
- Traceroute analysis
- Packet loss analysis
- Multi-port scanning
- Multi-target scanning
- Concurrent execution
- Configurable retries and timeouts
- Timing metrics for all checks
- Human-readable summaries
- JSON export mode
- File output support
- Quiet mode
- Docker support
- Exit codes for CI/CD automation
- IPv4/IPv6 support

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

# Docker

Build image:

```bash
docker build -t netcheckup .
```

Run container:

```bash
docker run --rm netcheckup google.com
```

Run with TLS:

```bash
docker run --rm netcheckup --tls google.com
```

Run traceroute:

```bash
docker run --rm netcheckup --traceroute google.com
```

Run packet loss analysis:

```bash
docker run --rm netcheckup --packet-loss google.com
```

Run port scan:

```bash
docker run --rm netcheckup --scan google.com
```

---

# Usage

```bash
netcheckup [options] <target> [target...]
```

---

# Examples

## Basic check

```bash
netcheckup google.com
```

## Multiple targets

```bash
netcheckup google.com github.com cloudflare.com
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

## Export JSON to file

```bash
netcheckup --json --output result.json google.com
```

## Quiet mode

```bash
netcheckup --quiet google.com
```

## TLS certificate inspection

```bash
netcheckup --tls google.com
```

## Traceroute

```bash
netcheckup --traceroute google.com
```

## Packet loss analysis

```bash
netcheckup --packet-loss google.com
```

## Custom DNS server

```bash
netcheckup --dns-server 1.1.1.1 google.com
```

## Port scanning

```bash
netcheckup --scan google.com
```

## Custom port scan

```bash
netcheckup --scan --ports 80,443,22 github.com
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
| `--tls` | Run TLS certificate inspection | `false` |
| `--traceroute` | Run traceroute analysis | `false` |
| `--packet-loss` | Run packet loss analysis | `false` |
| `--dns-server` | Use custom DNS resolver | `system default` |
| `--scan` | Run port scan | `false` |
| `--ports` | Ports to scan | `80,443,22` |
| `--output` | Write JSON output to file | `disabled` |
| `--version` | Show version information | `false` |

---

# Example Output

## Human-readable output

```text
========== SUMMARY ==========
✔ DNS resolution works (18 ms)
✔ Host reachable via ping (20 ms)
✔ TCP connection successful (31 ms)
✔ HTTP service responding (114 ms)

✔ TLS certificate valid (82 ms)
  Issuer: WR2
  Subject: *.google.com
  Expires in: 61 days

✔ Packet loss OK (4023 ms)
  Loss: 0%
  Packets: 5 sent / 5 received

✔ Port scan completed (88 ms)
  Open ports: 80, 443
  Closed ports: 22

→ Target is fully reachable and functioning
```

---

## JSON output

```json
{
  "target": "google.com",
  "primary_ip": "142.250.184.14",
  "dns_ok": true,
  "ping_ok": true,
  "tcp_ok": true,
  "http_ok": true,
  "dns_ms": 18,
  "ping_ms": 20,
  "tcp_ms": 31,
  "http_ms": 114,
  "tls": {
    "enabled": true,
    "ok": true,
    "issuer": "WR2",
    "subject": "*.google.com",
    "expires_at": "2026-07-14T08:35:44Z",
    "days_left": 61,
    "duration_ms": 82
  },
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
├── internal/
│   ├── checks/
│   │   ├── dns.go
│   │   ├── ping.go
│   │   ├── tcp.go
│   │   ├── http.go
│   │   ├── tls.go
│   │   ├── traceroute.go
│   │   ├── packetloss.go
│   │   ├── portscan.go
│   │   └── summary.go
│   └── version/
│       └── version.go
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

---

# Roadmap

- GitHub Actions CI/CD
- GitHub Releases
- Homebrew support
- CSV export
- HTML reports
- WHOIS lookup
- ASN lookup
- GeoIP lookup
- TUI dashboard
- Monitoring integrations

---

# License

MIT License
