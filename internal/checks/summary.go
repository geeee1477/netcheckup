package checks

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Target  string `json:"target"`
	DNS_OK  bool   `json:"dns_ok"`
	PING_OK bool   `json:"ping_ok"`
	TCP_OK  bool   `json:"tcp_ok"`
	HTTP_OK bool   `json:"http_ok"`
}

func PrintSummary(r Result) {
	fmt.Println("\n========== SUMMARY ==========")

	if r.DNS_OK {
		fmt.Println("✔ DNS resolution works")
	} else {
		fmt.Println("❌ DNS resolution failed")
	}

	if r.PING_OK {
		fmt.Println("✔ Host reachable via ping")
	} else {
		fmt.Println("❌ Host not reachable via ping")
	}

	if r.TCP_OK {
		fmt.Println("✔ TCP connection successful")
	} else {
		fmt.Println("❌ TCP connection failed")
	}

	if r.HTTP_OK {
		fmt.Println("✔ HTTP service responding")
	} else {
		fmt.Println("❌ HTTP request failed")
	}

	fmt.Println()
	fmt.Println(DiagnosisMessage(r))
}

func DiagnosisMessage(r Result) string {
	if !r.DNS_OK {
		return "→ Likely DNS or general connectivity issue"
	}

	if r.DNS_OK && !r.PING_OK && r.TCP_OK && r.HTTP_OK {
		return "→ Host blocks ICMP/ping, but TCP and HTTP are working"
	}

	if r.DNS_OK && !r.PING_OK && !r.TCP_OK && !r.HTTP_OK {
		return "→ Host not reachable at network level (routing, firewall, or offline)"
	}

	if r.DNS_OK && !r.PING_OK && r.TCP_OK && !r.HTTP_OK {
		return "→ ICMP is blocked and HTTP service may be failing"
	}

	if r.DNS_OK && r.PING_OK && !r.TCP_OK && !r.HTTP_OK {
		return "→ TCP port likely blocked or service down (no TCP + HTTP response)"
	}
	if r.DNS_OK && r.PING_OK && r.TCP_OK && !r.HTTP_OK {
		return "→ TCP port is reachable, but the application or HTTP service may be failing"
	}

	if r.DNS_OK && r.PING_OK && r.TCP_OK && r.HTTP_OK {
		return "→ Target is fully reachable and functioning"
	}

	return "→ No clear diagnosis available"
}

func PrintJSON(r Result) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	fmt.Println(string(data))
}
