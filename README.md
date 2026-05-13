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
