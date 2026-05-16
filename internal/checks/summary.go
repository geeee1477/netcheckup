package checks

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/geeee1477/netcheckup/internal/diagnosis"
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

	Diagnosis diagnosis.Diagnosis `json:"diagnosis"`
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

		if len(r.PortScan.OpenPorts) > 0 {
			fmt.Println("  Open ports:", strings.Join(r.PortScan.OpenPorts, ", "))
		}

		if len(r.PortScan.ClosedPorts) > 0 {
			fmt.Println("  Closed ports:", strings.Join(r.PortScan.ClosedPorts, ", "))
		}
	}

	PrintDiagnosisDetails(r)
}

func DiagnosisMessage(r Result) string {
	d := diagnosis.Analyze(diagnosis.CheckInput{
		DNSOK:  r.DNS_OK,
		PingOK: r.PING_OK,
		TCPOK:  r.TCP_OK,
		HTTPOK: r.HTTP_OK,
	})

	return "→ " + d.Summary
}

func DiagnosisObject(r Result) diagnosis.Diagnosis {
	return diagnosis.Analyze(diagnosis.CheckInput{
		DNSOK:  r.DNS_OK,
		PingOK: r.PING_OK,
		TCPOK:  r.TCP_OK,
		HTTPOK: r.HTTP_OK,
	})
}

func DiagnosisCode(r Result) string {
	d := diagnosis.Analyze(diagnosis.CheckInput{
		DNSOK:  r.DNS_OK,
		PingOK: r.PING_OK,
		TCPOK:  r.TCP_OK,
		HTTPOK: r.HTTP_OK,
	})

	return d.Code
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

func PrintDiagnosisDetails(r Result) {
	d := diagnosis.Analyze(diagnosis.CheckInput{
		DNSOK:  r.DNS_OK,
		PingOK: r.PING_OK,
		TCPOK:  r.TCP_OK,
		HTTPOK: r.HTTP_OK,
	})

	fmt.Println()
	fmt.Println("========== DIAGNOSIS ==========")
	fmt.Println("Severity:", d.Severity)
	fmt.Println("Summary:", d.Summary)
	fmt.Printf("Confidence: %d%%\n", d.Confidence)
	fmt.Println()
	fmt.Println("Explanation:")
	fmt.Println(d.Explanation)

	if len(d.Possible) > 0 {
		fmt.Println()
		fmt.Println("Likely causes:")
		for _, cause := range d.Possible {
			fmt.Printf("- %s (%d%%)\n", cause.Name, cause.Confidence)
		}
	} else if len(d.LikelyCauses) > 0 {
		fmt.Println()
		fmt.Println("Likely causes:")
		for _, cause := range d.LikelyCauses {
			fmt.Println("-", cause)
		}
	}

	if len(d.NextSteps) > 0 {
		fmt.Println()
		fmt.Println("Recommended next steps:")
		for i, step := range d.NextSteps {
			fmt.Printf("%d. %s\n", i+1, step)
		}
	}
}
