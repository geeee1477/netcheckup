package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/geeee1477/netcheckup/internal/checks"
)

func main() {
	port := flag.String("port", "443", "Port to check (default: 443)")
	jsonFlag := flag.Bool("json", false, "Output as JSON")

	flag.Usage = func() {
		fmt.Println("netcheckup - network diagnostic tool")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  netcheckup [--port <port>] [--json] <target>")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  netcheckup google.com")
		fmt.Println("  netcheckup --port 80 google.com")
		fmt.Println("  netcheckup --json google.com")
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	target := flag.Arg(0)

	if !*jsonFlag {
		fmt.Println("netcheckup starting...\n")
	}

	dnsOK := checks.ResolveDNS(target)
	pingOK := checks.CheckPing(target)
	tcpOK := checks.CheckTCP(target, *port)
	httpOK := checks.CheckHTTP(target, *port)

	result := checks.Result{
		Target:  target,
		DNS_OK:  dnsOK,
		PING_OK: pingOK,
		TCP_OK:  tcpOK,
		HTTP_OK: httpOK,
	}

	if *jsonFlag {
		checks.PrintJSON(result)
		return
	}

	checks.PrintSummary(result)
}
