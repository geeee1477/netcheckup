package checks

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type PortScanResult struct {
	Enabled     bool     `json:"enabled"`
	OpenPorts   []string `json:"open_ports,omitempty"`
	ClosedPorts []string `json:"closed_ports,omitempty"`
	DurationMS  int64    `json:"duration_ms"`
}

func ScanPorts(
	target string,
	ports string,
	timeout int,
	verbose bool,
) PortScanResult {

	start := time.Now()

	result := PortScanResult{
		Enabled: true,
	}

	if verbose {
		fmt.Println()
		fmt.Println("[SCAN] Starting port scan:", target)
	}

	portList := strings.Split(ports, ",")

	for _, port := range portList {

		port = strings.TrimSpace(port)

		address := net.JoinHostPort(target, port)

		conn, err := net.DialTimeout(
			"tcp",
			address,
			time.Duration(timeout)*time.Second,
		)

		if err != nil {

			result.ClosedPorts = append(result.ClosedPorts, port)

			if verbose {
				fmt.Println("[SCAN] ❌", port, "closed")
			}

			continue
		}

		conn.Close()

		result.OpenPorts = append(result.OpenPorts, port)

		if verbose {
			fmt.Println("[SCAN] ✅", port, "open")
		}
	}

	result.DurationMS = time.Since(start).Milliseconds()

	return result
}
