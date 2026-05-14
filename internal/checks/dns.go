package checks

import (
	"context"
	"fmt"
	"net"
	"time"
)

func ResolveDNS(target string, dnsServer string, timeout, retries int, verbose bool) (bool, string, int64) {
	start := time.Now()

	if verbose {
		fmt.Println("[DNS] Checking:", target)
	}

	resolver := &net.Resolver{}

	if dnsServer != "" {
		resolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{
					Timeout: time.Duration(timeout) * time.Second,
				}

				return d.DialContext(ctx, "udp", dnsServer+":53")
			},
		}
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Second,
	)

	defer cancel()

	ips, err := resolver.LookupHost(ctx, target)

	duration := time.Since(start).Milliseconds()

	if err != nil {
		if verbose {
			fmt.Println("[DNS] ❌ Resolution failed")
			fmt.Println("Error:", err)
		}

		return false, "", duration
	}

	if len(ips) == 0 {
		if verbose {
			fmt.Println("[DNS] ❌ No IPs returned")
		}

		return false, "", duration
	}

	primaryIP := ips[0]

	if verbose {
		fmt.Println("[DNS] ✅ Resolution successful")
		fmt.Println("Resolved IPs:", len(ips), "found")
		fmt.Println("Primary IP:", primaryIP)

		if dnsServer != "" {
			fmt.Println("DNS Server:", dnsServer)
		}
	}

	return true, primaryIP, duration
}
