package checks

import (
	"fmt"
	"os/exec"
	"strings"
)

func CheckPing(target string) bool {
	fmt.Println("\n[PING] Checking:", target)

	// macOS/Linux: -c 1 = 1 Paket
	cmd := exec.Command("ping", "-c", "1", target)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("[PING] ❌ Failed")
		fmt.Println("→ Possible causes:")
		fmt.Println(" - host unreachable")
		fmt.Println(" - ICMP blocked by firewall")
		fmt.Println(" - network issue")
		fmt.Println("Error:", err)
		return false
	}

	out := string(output)

	// einfache Auswertung
	if strings.Contains(out, "1 packets received") || strings.Contains(out, "1 received") {
		fmt.Println("[PING] ✅ Host reachable")
		return true
	}

	fmt.Println("[PING] ⚠️ No response")
	return false
}
