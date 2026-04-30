package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/geeee1477/netcheckup/internal/checks"
)

func main() {
	port := flag.String("port", "443", "Port to check (default: 443)")
	flag.Usage = func() {
		fmt.Println("netcheckup - network diagnostic tool")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  netcheckup [--port <port>] <target>")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  netcheckup google.com")
		fmt.Println("  netcheckup --port 80 google.com")
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	target := flag.Arg(0)

	fmt.Println("netcheckup starting...\n")

	dnsOK := checks.ResolveDNS(target)
	pingOK := checks.CheckPing(target)
	tcpOK := checks.CheckTCP(target, *port)
	httpOK := checks.CheckHTTP(target, *port)

	result := checks.Result{
		DNS_OK:  dnsOK,
		PING_OK: pingOK,
		TCP_OK:  tcpOK,
		HTTP_OK: httpOK,
	}

	checks.PrintSummary(result)
}
