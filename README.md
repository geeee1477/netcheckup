# netcheckup

A lightweight network diagnostics CLI tool to analyze connectivity issues and guide users towards possible root causes.

---

## 🚀 Features

- DNS resolution check  
- TCP connectivity check (custom port)  
- HTTP/HTTPS request check  
- Automatic diagnostic summary with interpretation  

---

## 🧠 How it works

The tool follows a structured troubleshooting flow:

1. DNS → resolves domain to IP  
2. TCP → checks if the port is reachable  
3. HTTP → verifies if the service responds  
4. Summary → interprets the results  

---

## 📦 Usage

```bash
go run cmd/netcheckup/main.go <target>
```

Example:

```bash
go run cmd/netcheckup/main.go google.com
```

Custom port:

```bash
go run cmd/netcheckup/main.go google.com --port 80
```

---

## 🧪 Example Output

```text
[DNS] ✅ Resolution successful
[TCP] ✅ Port 443 is reachable
[HTTP] ✅ Response received
Status: 200 OK

========== SUMMARY ==========
✔ DNS resolution works
✔ TCP connection successful
✔ HTTP service responding

→ Target is fully reachable and functioning
```

---

## 🛠 Tech Stack

- Go (Golang)  
- Standard library only (no external dependencies)  

---

## 📈 Roadmap

- [ ] Ping / ICMP check  
- [ ] Traceroute  
- [ ] JSON output mode  
- [ ] Better CLI (flags & help)  
- [ ] Parallel checks  
- [ ] Logging & debug mode  

---

## 🎯 Goal

Build a practical tool that reflects real-world network troubleshooting and can be used in IT support / system engineering contexts.
