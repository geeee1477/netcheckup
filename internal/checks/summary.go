package checks

import "fmt"

type Result struct {
	DNS_OK  bool
	PING_OK bool
	TCP_OK  bool
	HTTP_OK bool
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

	if !r.DNS_OK {
		fmt.Println("→ Likely DNS or general connectivity issue")
		return
	}

	if r.DNS_OK && !r.PING_OK && r.TCP_OK && r.HTTP_OK {
		fmt.Println("→ Host blocks ICMP/ping, but TCP and HTTP are working")
		return
	}

	if r.DNS_OK && !r.PING_OK && !r.TCP_OK {
		fmt.Println("→ Likely network, routing, or firewall issue")
		return
	}

	if r.DNS_OK && r.PING_OK && !r.TCP_OK {
		fmt.Println("→ Host is reachable, but the selected TCP port may be blocked or closed")
		return
	}

	if r.TCP_OK && !r.HTTP_OK {
		fmt.Println("→ TCP port is reachable, but the application or HTTP service may be failing")
		return
	}

	if r.DNS_OK && r.PING_OK && r.TCP_OK && r.HTTP_OK {
		fmt.Println("→ Target is fully reachable and functioning")
	}
}
