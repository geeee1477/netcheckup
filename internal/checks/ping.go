package checks

import (
	"fmt"
	"os/exec"
	"time"
)

func CheckPing(target string, timeout, retries int, verbose bool) (bool, int64) {
	start := time.Now()

	if verbose {
		fmt.Println("\n[PING] Checking:", target)
	}

	for attempt := 1; attempt <= retries; attempt++ {
		cmd := exec.Command("ping", "-c", "1", "-W", fmt.Sprintf("%d", timeout), target)

		err := cmd.Run()
		if err == nil {
			if verbose {
				fmt.Println("[PING] ✅ Host reachable")
			}
			return true, time.Since(start).Milliseconds()
		}

		if verbose {
			fmt.Printf("[PING] Attempt %d/%d failed\n", attempt, retries)
			fmt.Println("Error:", err)
		}
	}

	if verbose {
		fmt.Println("[PING] ❌ Failed")
		fmt.Println("→ Possible causes:")
		fmt.Println(" - host unreachable")
		fmt.Println(" - ICMP blocked by firewall")
		fmt.Println(" - network issue")
	}

	return false, time.Since(start).Milliseconds()
}
