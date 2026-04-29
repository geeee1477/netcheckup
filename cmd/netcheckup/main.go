package main

import (
	"fmt"
	"os"

	"github.com/geeee1477/netcheckup/internal/checks"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: netcheckup <target>")
		return
	}

	target := os.Args[1]

	fmt.Println("netcheckup starting...")

	checks.ResolveDNS(target)
	checks.CheckTCP(target, "443")
}
