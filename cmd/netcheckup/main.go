package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/geeee1477/netcheckup/internal/checks"
)

func main() {
	port := flag.String("port", "443", "Port to check (default: 443)")
	timeout := flag.Int("timeout", 3, "Timeout in seconds (default: 3)")
	retries := flag.Int("retries", 2, "Number of retries (default: 2)")
	jsonFlag := flag.Bool("json", false, "Output as JSON")

	flag.Usage = func() {
		fmt.Println("netcheckup - network diagnostic tool")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  netcheckup [--port <port>] [--timeout <seconds>] [--retries <n>] [--json] <target>")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  netcheckup google.com")
		fmt.Println("  netcheckup --port 80 google.com")
		fmt.Println("  netcheckup --timeout 2 google.com")
		fmt.Println("  netcheckup --retries 3 google.com")
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

	var (
		dnsOK, pingOK, tcpOK, httpOK bool
		primaryIP                    string
		wg                           sync.WaitGroup
	)

	wg.Add(4)

	// DNS
	go func() {
		defer wg.Done()
		dnsOK, primaryIP = checks.ResolveDNS(target, verbose)
	}()

	// PING
	go func() {
		defer wg.Done()
		pingOK = checks.CheckPing(target, verbose)
	}()

	// TCP
	go func() {
		defer wg.Done()
		tcpOK = checks.CheckTCP(target, *port, *timeout, *retries, verbose)
	}()

	// HTTP
	go func() {
		defer wg.Done()
		httpOK = checks.CheckHTTP(target, *port, *timeout, *retries, verbose)
	}()

	wg.Wait()

	result := checks.Result{
		Target:    target,
		PrimaryIP: primaryIP,
		DNS_OK:    dnsOK,
		PING_OK:   pingOK,
		TCP_OK:    tcpOK,
		HTTP_OK:   httpOK,
	}

	if *jsonFlag {
		result.Diagnosis = checks.DiagnosisCode(result)
		checks.PrintJSON(result)
		return
	}

	result.Diagnosis = checks.DiagnosisMessage(result)
	checks.PrintSummary(result)
}
