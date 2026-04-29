package checks

import "fmt"

type Result struct {
	DNS_OK  bool
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
		fmt.Println("→ Likely DNS or connectivity issue")
		return
	}

	if r.DNS_OK && !r.TCP_OK {
		fmt.Println("→ Likely firewall or network issue")
		return
	}

	if r.TCP_OK && !r.HTTP_OK {
		fmt.Println("→ Service reachable but application may be down")
		return
	}

	if r.DNS_OK && r.TCP_OK && r.HTTP_OK {
		fmt.Println("→ Target is fully reachable and functioning")
	}
}
