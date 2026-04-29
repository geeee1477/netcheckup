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
	port := "443" // default

	// simple flag parsing
	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] == "--port" && i+1 < len(os.Args) {
			port = os.Args[i+1]
		}
	}

	// remove accidental port in target
	if strings.Contains(target, ":") {
		parts := strings.Split(target, ":")
		target = parts[0]
	}

	fmt.Println("netcheckup starting...")

	checks.ResolveDNS(target)
	checks.CheckTCP(target, port)
	checks.CheckHTTP(target, port)
}
