package checks

import (
	"fmt"
	"net"
	"time"
)

func CheckTCP(target string, port string, verbose bool) bool {
	address := target + ":" + port

	if verbose {
		fmt.Println("\n[TCP] Checking:", address)
	}

	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		if verbose {
			fmt.Println("[TCP] ❌ Connection failed")
			fmt.Println("→ Possible causes:")
			fmt.Println(" - firewall blocking the port")
			fmt.Println(" - service not running")
			fmt.Println(" - network connectivity issue")
			fmt.Println("Error:", err)
		}
		return false
	}

	conn.Close()

	if verbose {
		fmt.Println("[TCP] ✅ Port", port, "is reachable")
	}

	return true
}
