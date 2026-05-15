package checks

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Result struct {
	Target    string `json:"target"`
	PrimaryIP string `json:"primary_ip,omitempty"`

	DNS_OK  bool `json:"dns_ok"`
	PING_OK bool `json:"ping_ok"`
	TCP_OK  bool `json:"tcp_ok"`
	HTTP_OK bool `json:"http_ok"`

	DNS_MS     int64            `json:"dns_ms,omitempty"`
	PING_MS    int64            `json:"ping_ms,omitempty"`
	TCP_MS     int64            `json:"tcp_ms,omitempty"`
	HTTP_MS    int64            `json:"http_ms,omitempty"`
	TLS        TLSResult        `json:"tls,omitempty"`
	Traceroute TracerouteResult `json:"traceroute,omitempty"`
	PacketLoss PacketLossResult `json:"packet_loss,omitempty"`
	PortScan   PortScanResult   `json:"port_scan,omitempty"`

	Diagnosis string `json:"diagnosis"`
}

func PrintSummary(r Result) {
	fmt.Println("\n========== SUMMARY ==========")

	if r.DNS_OK {
		fmt.Printf("✔ DNS resolution works (%d ms)\n", r.DNS_MS)
	} else {
		fmt.Printf("❌ DNS resolution failed (%d ms)\n", r.DNS_MS)
	}

	if r.PING_OK {
		fmt.Printf("✔ Host reachable via ping (%d ms)\n", r.PING_MS)
	} else {
		fmt.Printf("❌ Host not reachable via ping (%d ms)\n", r.PING_MS)
	}

	if r.TCP_OK {
		fmt.Printf("✔ TCP connection successful (%d ms)\n", r.TCP_MS)
	} else {
		fmt.Printf("❌ TCP connection failed (%d ms)\n", r.TCP_MS)
	}

	if r.HTTP_OK {
		fmt.Printf("✔ HTTP service responding (%d ms)\n", r.HTTP_MS)
	} else {
		fmt.Printf("❌ HTTP request failed (%d ms)\n", r.HTTP_MS)
	}

	if r.TLS.Enabled {
		fmt.Println()

		if r.TLS.OK {
			fmt.Printf("✔ TLS certificate valid (%d ms)\n", r.TLS.DurationMS)
			fmt.Println("  Issuer:", r.TLS.Issuer)
			fmt.Println("  Subject:", r.TLS.Subject)
			fmt.Println("  Expires in:", r.TLS.DaysLeft, "days")
		} else {
			fmt.Printf("❌ TLS certificate invalid (%d ms)\n", r.TLS.DurationMS)
			fmt.Println("  Error:", r.TLS.Error)
		}
	}

	if r.Traceroute.Enabled {
		fmt.Println()

		if r.Traceroute.OK {
			fmt.Printf("✔ Traceroute completed (%d ms)\n", r.Traceroute.DurationMS)
			fmt.Println("  Hops:", len(r.Traceroute.Hops))
		} else {
			fmt.Printf("❌ Traceroute failed (%d ms)\n", r.Traceroute.DurationMS)
			fmt.Println("  Error:", r.Traceroute.Error)
		}
	}

	if r.PacketLoss.Enabled {
		fmt.Println()

		if r.PacketLoss.OK {
			fmt.Printf("✔ Packet loss OK (%d ms)\n", r.PacketLoss.DurationMS)
		} else {
			fmt.Printf("❌ Packet loss detected (%d ms)\n", r.PacketLoss.DurationMS)
		}

		fmt.Printf("  Loss: %d%%\n", r.PacketLoss.LossPercent)
		fmt.Printf("  Packets: %d sent / %d received\n",
			r.PacketLoss.Sent,
			r.PacketLoss.Received,
		)
	}

	if r.PortScan.Enabled {
		fmt.Println()

		fmt.Printf("✔ Port scan completed (%d ms)\n", r.PortScan.DurationMS)

		fmt.Println("  Open ports:", strings.Join(r.PortScan.OpenPorts, ", "))

		if len(r.PortScan.ClosedPorts) > 0 {
			fmt.Println("  Closed ports:", strings.Join(r.PortScan.ClosedPorts, ", "))
		}
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

func DiagnosisCode(r Result) string {
	if !r.DNS_OK {
		return "dns_or_connectivity_issue"
	}

	if r.DNS_OK && !r.PING_OK && r.TCP_OK && r.HTTP_OK {
		return "icmp_blocked_service_reachable"
	}

	if r.DNS_OK && !r.PING_OK && !r.TCP_OK && !r.HTTP_OK {
		return "network_routing_firewall_or_offline"
	}

	if r.DNS_OK && !r.PING_OK && r.TCP_OK && !r.HTTP_OK {
		return "icmp_blocked_http_failing"
	}

	if r.DNS_OK && r.PING_OK && !r.TCP_OK && !r.HTTP_OK {
		return "tcp_port_blocked_or_service_down"
	}

	if r.DNS_OK && r.PING_OK && r.TCP_OK && !r.HTTP_OK {
		return "http_service_failing"
	}

	if r.DNS_OK && r.PING_OK && r.TCP_OK && r.HTTP_OK {
		return "target_fully_reachable"
	}

	return "unknown"
}

func PrintJSON(r Result) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}

	fmt.Println(string(data))
}

func JSONString(r Result) (string, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
