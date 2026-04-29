package checks

import (
	"fmt"
	"os/exec"
)

func CheckPing(target string) bool {
	fmt.Println("\n[PING] Checking:", target)

	cmd := exec.Command("ping", "-c", "1", "-W", "1", target)

	err := cmd.Run()
	if err != nil {
		fmt.Println("[PING] ❌ Failed")
		fmt.Println("→ Possible causes:")
		fmt.Println(" - host unreachable")
		fmt.Println(" - ICMP blocked by firewall")
		fmt.Println(" - network issue")
		fmt.Println("Error:", err)
		return false
	}

	fmt.Println("[PING] ✅ Host reachable")
	return true
}
