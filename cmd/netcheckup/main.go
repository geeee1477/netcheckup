package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/geeee1477/netcheckup/internal/checks"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: netcheckup <target> [--port <port>]")
		return
	}

	target := os.Args[1]
	port := "443"

	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] == "--port" && i+1 < len(os.Args) {
			port = os.Args[i+1]
		}
	}

	if strings.Contains(target, ":") {
		parts := strings.Split(target, ":")
		target = parts[0]
	}

	fmt.Println("netcheckup starting...\n")

	dnsOK := checks.ResolveDNS(target)
	tcpOK := checks.CheckTCP(target, port)
	httpOK := checks.CheckHTTP(target, port)

	result := checks.Result{
		DNS_OK:  dnsOK,
		TCP_OK:  tcpOK,
		HTTP_OK: httpOK,
	}

	checks.PrintSummary(result)
}
