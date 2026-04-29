package checks

import (
	"fmt"
	"net"
	"time"
)

func CheckTCP(target string, port string) {
	address := target + ":" + port

	fmt.Println("\n[TCP] Checking:", address)

	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		fmt.Println("[TCP] ❌ Connection failed")
		fmt.Println("→ Possible causes:")
		fmt.Println(" - firewall blocking the port")
		fmt.Println(" - service not running")
		fmt.Println(" - network connectivity issue")
		fmt.Println("Error:", err)
		return
	}

	conn.Close()

	fmt.Println("[TCP] ✅ Port", port, "is reachable")
}
