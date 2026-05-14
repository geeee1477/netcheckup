package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/geeee1477/netcheckup/internal/checks"
	"github.com/geeee1477/netcheckup/internal/version"
)

func main() {
	port := flag.String("port", "443", "Port to check (default: 443)")
	timeout := flag.Int("timeout", 3, "Timeout in seconds (default: 3)")
	retries := flag.Int("retries", 2, "Number of retries (default: 2)")
	jsonFlag := flag.Bool("json", false, "Output as JSON")
	tlsFlag := flag.Bool("tls", false, "Run TLS certificate inspection")
	tracerouteFlag := flag.Bool("traceroute", false, "Run traceroute")
	packetLossFlag := flag.Bool("packet-loss", false, "Run packet loss analysis")
	versionFlag := flag.Bool("version", false, "Show version")
	quietFlag := flag.Bool("quiet", false, "Hide verbose logs")

	flag.Usage = func() {
		fmt.Println("netcheckup - network diagnostic tool")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  netcheckup [--port <port>] [--timeout <seconds>] [--retries <n>] [--json] <target>")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  netcheckup --tls google.com")
		fmt.Println("  netcheckup google.com")
		fmt.Println("  netcheckup --port 80 google.com")
		fmt.Println("  netcheckup --timeout 2 google.com")
		fmt.Println("  netcheckup --retries 3 google.com")
		fmt.Println("  netcheckup --json google.com")
		fmt.Println("  netcheckup --traceroute google.com")
		fmt.Println("  netcheckup --packet-loss google.com")

	}

	flag.Parse()

	if *versionFlag {
		fmt.Println("netcheckup version", version.Version)
		return
	}

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	target := flag.Arg(0)
	verbose := !*jsonFlag && !*quietFlag

	if verbose {
		fmt.Println("netcheckup starting...\n")
	}

	var (
		dnsOK, pingOK, tcpOK, httpOK bool
		primaryIP                    string
		dnsMS, pingMS, tcpMS, httpMS int64
		tlsResult                    checks.TLSResult
		tracerouteResult             checks.TracerouteResult
		packetLossResult             checks.PacketLossResult
		wg                           sync.WaitGroup
	)

	wg.Add(4)

	go func() {
		defer wg.Done()
		dnsOK, primaryIP, dnsMS = checks.ResolveDNS(target, *timeout, *retries, verbose)
	}()

	go func() {
		defer wg.Done()
		pingOK, pingMS = checks.CheckPing(target, *timeout, *retries, verbose)
	}()

	go func() {
		defer wg.Done()
		tcpOK, tcpMS = checks.CheckTCP(target, *port, *timeout, *retries, verbose)
	}()

	go func() {
		defer wg.Done()
		httpOK, httpMS = checks.CheckHTTP(target, *port, *timeout, *retries, verbose)
	}()

	wg.Wait()

	if *tlsFlag {
		tlsResult = checks.CheckTLS(target, *port, *timeout, verbose)
	}

	if *tracerouteFlag {
		tracerouteResult = checks.RunTraceroute(target, *timeout, verbose)
	}

	if *packetLossFlag {
		packetLossResult = checks.CheckPacketLoss(target, 5, *timeout, verbose)
	}

	result := checks.Result{
		Target:     target,
		PrimaryIP:  primaryIP,
		DNS_OK:     dnsOK,
		PING_OK:    pingOK,
		TCP_OK:     tcpOK,
		HTTP_OK:    httpOK,
		DNS_MS:     dnsMS,
		PING_MS:    pingMS,
		TCP_MS:     tcpMS,
		HTTP_MS:    httpMS,
		TLS:        tlsResult,
		Traceroute: tracerouteResult,
		PacketLoss: packetLossResult,
	}

	if *jsonFlag {
		result.Diagnosis = checks.DiagnosisCode(result)
		checks.PrintJSON(result)

		if !result.DNS_OK || !result.PING_OK || !result.TCP_OK || !result.HTTP_OK {
			os.Exit(1)
		}

		return
	}

	result.Diagnosis = checks.DiagnosisMessage(result)
	checks.PrintSummary(result)

	if !result.DNS_OK || !result.PING_OK || !result.TCP_OK || !result.HTTP_OK {
		os.Exit(1)
	}
}
