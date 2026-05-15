package diagnosis

type CheckInput struct {
	DNSOK  bool
	PingOK bool
	TCPOK  bool
	HTTPOK bool
}

func Analyze(input CheckInput) Diagnosis {
	if !input.DNSOK {
		return Diagnosis{
			Code:        "dns_resolution_failed",
			Severity:    "critical",
			Summary:     "DNS resolution failed",
			Explanation: "The target could not be resolved to an IP address. Without DNS resolution, TCP and HTTP checks cannot reliably reach the intended host.",
			LikelyCauses: []string{
				"DNS server unreachable",
				"Domain name typo",
				"VPN or split-DNS issue",
				"Local resolver misconfiguration",
			},
			NextSteps: []string{
				"Try a custom DNS server: --dns-server 1.1.1.1",
				"Check whether the domain name is spelled correctly.",
				"Verify local DNS, VPN, or firewall settings.",
			},
		}
	}

	if input.DNSOK && !input.PingOK && input.TCPOK && input.HTTPOK {
		return Diagnosis{
			Code:        "icmp_blocked_service_reachable",
			Severity:    "info",
			Summary:     "ICMP appears to be blocked, but the service is reachable",
			Explanation: "DNS resolution works. Ping failed, but TCP and HTTP checks succeeded. This usually means the host or firewall blocks ICMP while the actual service is healthy.",
			LikelyCauses: []string{
				"ICMP blocked by firewall",
				"Cloud provider blocks ping",
				"Host configured to ignore ping",
			},
			NextSteps: []string{
				"Do not treat failed ping alone as an outage.",
				"Use TCP and HTTP checks to verify service availability.",
				"Check firewall or security group rules if ICMP is required.",
			},
		}
	}

	if input.DNSOK && input.PingOK && !input.TCPOK && !input.HTTPOK {
		return Diagnosis{
			Code:        "tcp_port_blocked_or_service_down",
			Severity:    "critical",
			Summary:     "Target is reachable, but the TCP service is not responding",
			Explanation: "DNS resolution works and the host responds to ping, but TCP and HTTP checks failed. This indicates that the target is reachable at network level, but the selected port may be blocked or the service may be down.",
			LikelyCauses: []string{
				"Service not running",
				"TCP port blocked by firewall",
				"Wrong port selected",
				"NAT or security group issue",
			},
			NextSteps: []string{
				"Verify that the service is running on the target port.",
				"Check local firewall, server firewall, cloud security groups, or NAT rules.",
				"Try a port scan: --scan --ports 80,443,22",
			},
		}
	}

	if input.DNSOK && input.PingOK && input.TCPOK && !input.HTTPOK {
		return Diagnosis{
			Code:        "http_service_failing",
			Severity:    "warning",
			Summary:     "TCP is reachable, but HTTP is not responding correctly",
			Explanation: "DNS resolution works, the host is reachable, and the TCP port is open. HTTP failed, which usually points to an application, TLS, reverse proxy, or web server issue.",
			LikelyCauses: []string{
				"Web server issue",
				"Reverse proxy misconfiguration",
				"TLS or certificate problem",
				"Application not responding",
			},
			NextSteps: []string{
				"Check the web server or reverse proxy logs.",
				"Run TLS inspection with: --tls",
				"Verify whether the correct protocol and port are being used.",
			},
		}
	}

	if input.DNSOK && !input.PingOK && !input.TCPOK && !input.HTTPOK {
		return Diagnosis{
			Code:        "network_routing_firewall_or_offline",
			Severity:    "critical",
			Summary:     "Target does not appear reachable beyond DNS",
			Explanation: "DNS resolution works, so the hostname exists. Ping, TCP, and HTTP checks all failed. This may indicate routing issues, firewall filtering, VPN problems, or an offline host.",
			LikelyCauses: []string{
				"Host offline",
				"Routing issue",
				"Firewall filtering",
				"VPN connectivity issue",
			},
			NextSteps: []string{
				"Run traceroute with: --traceroute",
				"Test from another network or VPN.",
				"Check whether the target host is online and reachable from your network.",
			},
		}
	}

	if input.DNSOK && input.PingOK && input.TCPOK && input.HTTPOK {
		return Diagnosis{
			Code:         "target_fully_reachable",
			Severity:     "ok",
			Summary:      "Target is fully reachable and functioning",
			Explanation:  "DNS resolution works, the host responds to ping, the TCP port is reachable, and HTTP responded successfully.",
			LikelyCauses: []string{},
			NextSteps: []string{
				"No immediate network issue detected.",
				"If users still report problems, check authentication, application logs, browser/client issues, or regional routing.",
			},
		}
	}

	return Diagnosis{
		Code:        "unknown",
		Severity:    "unknown",
		Summary:     "No clear diagnosis available",
		Explanation: "The check results do not match a known diagnosis pattern yet.",
		LikelyCauses: []string{
			"Mixed or uncommon failure pattern",
		},
		NextSteps: []string{
			"Run a full diagnostic with TLS, traceroute, packet loss, and port scan.",
			"Compare the result from another network.",
		},
	}
}
