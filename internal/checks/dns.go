package checks

import (
	"fmt"
	"net"
)

func ResolveDNS(target string, verbose bool) (bool, string) {
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
		return false, ""
	}

	if len(ips) == 0 {
		if verbose {
			fmt.Println("[DNS] ⚠️ No IPs returned")
			fmt.Println("→ Possible causes:")
			fmt.Println(" - DNS misconfiguration")
			fmt.Println(" - domain has no A/AAAA records")
		}
		return false, ""
	}

	primaryIP := ""
	for _, ip := range ips {
		parsed := net.ParseIP(ip)
		if parsed != nil && parsed.To4() != nil && !parsed.IsPrivate() {
			primaryIP = ip
			break
		}
	}

	if primaryIP == "" {
		primaryIP = ips[0]
	}

	if verbose {
		fmt.Println("[DNS] ✅ Resolution successful")
		fmt.Println("→ DNS is working correctly")
		fmt.Println("Resolved IPs:", len(ips), "found")
		fmt.Println("Primary IP:", primaryIP)
	}

	return true, primaryIP
}
