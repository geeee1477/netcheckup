package checks

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type TracerouteResult struct {
	Enabled    bool     `json:"enabled"`
	OK         bool     `json:"ok"`
	Hops       []string `json:"hops,omitempty"`
	DurationMS int64    `json:"duration_ms"`
	Error      string   `json:"error,omitempty"`
}

func RunTraceroute(target string, timeout int, verbose bool) TracerouteResult {
	start := time.Now()

	result := TracerouteResult{
		Enabled: true,
	}

	if verbose {
		fmt.Println()
		fmt.Println("[TRACE] Running traceroute:", target)
	}

	cmd := exec.Command(
		"traceroute",
		"-m", "12",
		"-w", fmt.Sprintf("%d", timeout),
		target,
	)

	output, err := cmd.CombinedOutput()

	result.DurationMS = time.Since(start).Milliseconds()

	if err != nil {
		result.OK = false
		result.Error = err.Error()

		if verbose {
			fmt.Println("[TRACE] ❌ Traceroute failed")
			fmt.Println("Error:", err)
		}

		return result
	}

	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		result.Hops = append(result.Hops, line)
	}

	result.OK = true

	if verbose {
		fmt.Println("[TRACE] ✅ Traceroute completed")
		fmt.Println("[TRACE] Hops:", len(result.Hops))
	}

	return result
}
