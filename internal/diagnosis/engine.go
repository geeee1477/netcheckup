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
			Confidence:  85,
			LikelyCauses: []string{
				"DNS server unreachable",
				"Domain name typo",
				"VPN or split-DNS issue",
				"Local resolver misconfiguration",
			},
			Possible: []PossibleCause{
				{Name: "dns_server_unreachable", Confidence: 85},
				{Name: "domain_name_typo", Confidence: 70},
				{Name: "vpn_or_split_dns_issue", Confidence: 65},
				{Name: "local_resolver_misconfiguration", Confidence: 60},
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
			Confidence:  90,
			LikelyCauses: []string{
				"ICMP blocked by firewall",
				"Cloud provider blocks ping",
				"Host configured to ignore ping",
			},
			Possible: []PossibleCause{
				{Name: "icmp_blocked_by_firewall", Confidence: 90},
				{Name: "cloud_provider_blocks_ping", Confidence: 75},
				{Name: "host_ignores_ping", Confidence: 70},
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
			Confidence:  82,
			LikelyCauses: []string{
				"Service not running",
				"TCP port blocked by firewall",
				"Wrong port selected",
				"NAT or security group issue",
			},
			Possible: []PossibleCause{
				{Name: "tcp_port_blocked_by_firewall", Confidence: 82},
				{Name: "service_not_running", Confidence: 78},
				{Name: "wrong_port_selected", Confidence: 55},
				{Name: "nat_or_security_group_issue", Confidence: 65},
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
			Confidence:  80,
			LikelyCauses: []string{
				"Web server issue",
				"Reverse proxy misconfiguration",
				"TLS or certificate problem",
				"Application not responding",
			},
			Possible: []PossibleCause{
				{Name: "reverse_proxy_misconfiguration", Confidence: 80},
				{Name: "web_server_issue", Confidence: 75},
				{Name: "tls_or_certificate_problem", Confidence: 65},
				{Name: "application_not_responding", Confidence: 70},
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
			Confidence:  78,
			LikelyCauses: []string{
				"Host offline",
				"Routing issue",
				"Firewall filtering",
				"VPN connectivity issue",
			},
			Possible: []PossibleCause{
				{Name: "firewall_filtering", Confidence: 78},
				{Name: "host_offline", Confidence: 72},
				{Name: "routing_issue", Confidence: 68},
				{Name: "vpn_connectivity_issue", Confidence: 60},
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
			Code:        "target_fully_reachable",
			Severity:    "ok",
			Summary:     "Target is fully reachable and functioning",
			Explanation: "DNS resolution works, the host responds to ping, the TCP port is reachable, and HTTP responded successfully.",
			Confidence:  95,
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
		Confidence:  40,
		LikelyCauses: []string{
			"Mixed or uncommon failure pattern",
		},
		Possible: []PossibleCause{
			{Name: "mixed_or_uncommon_failure_pattern", Confidence: 40},
		},
		NextSteps: []string{
			"Run a full diagnostic with TLS, traceroute, packet loss, and port scan.",
			"Compare the result from another network.",
		},
	}
}
