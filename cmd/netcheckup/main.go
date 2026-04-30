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
	verbose := !*jsonFlag

	if verbose {
		fmt.Println("netcheckup starting...\n")
	}

	var dnsOK, pingOK, tcpOK, httpOK bool

	done := make(chan bool)

	// DNS
	go func() {
		dnsOK = checks.ResolveDNS(target, verbose)
		done <- true
	}()

	// PING
	go func() {
		pingOK = checks.CheckPing(target, verbose)
		done <- true
	}()

	// TCP
	go func() {
		tcpOK = checks.CheckTCP(target, *port, verbose)
		done <- true
	}()

	// HTTP
	go func() {
		httpOK = checks.CheckHTTP(target, *port, verbose)
		done <- true
	}()

	// warten bis alle fertig sind
	for i := 0; i < 4; i++ {
		<-done
	}
	
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
