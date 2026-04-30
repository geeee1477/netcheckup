package checks

import (
	"fmt"
	"net"
)

func ResolveDNS(target string, verbose bool) bool {
	if verbose {
		fmt.Println("[DNS] Checking:", target)
	}

	ips, err := net.LookupHost(target)
	if err != nil {
		if verbose {
			fmt.Println("[DNS] ❌ Resolution failed")
			fmt.Println("→ Possible causes:")
			fmt.Println(" - DNS server unreachable")
			fmt.Println(" - no internet connectivity")
			fmt.Println(" - misconfigured resolver")
			fmt.Println("Error:", err)
		}
		return false
	}

	if len(ips) == 0 {
		if verbose {
			fmt.Println("[DNS] ⚠️ No IPs returned")
			fmt.Println("→ Possible causes:")
			fmt.Println(" - DNS misconfiguration")
			fmt.Println(" - domain has no A/AAAA records")
		}
		return false
	}

	if verbose {
		fmt.Println("[DNS] ✅ Resolution successful")
		fmt.Println("→ DNS is working correctly")
		fmt.Println("Resolved IPs:", len(ips), "found")
		fmt.Println("Primary IP:", ips[0])
	}

	return true
}
