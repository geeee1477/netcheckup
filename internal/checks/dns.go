package checks

import (
	"fmt"
	"net"
)

func ResolveDNS(target string) {
	fmt.Println("[DNS] Checking:", target)

	ips, err := net.LookupHost(target)
	if err != nil {
		fmt.Println("[DNS] ❌ Failed:", err)
		return
	}

	fmt.Println("[DNS] ✅ Resolved IPs:")
	for _, ip := range ips {
		fmt.Println(" -", ip)
	}
}
