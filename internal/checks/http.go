package checks

import (
	"fmt"
	"net/http"
	"time"
)

func CheckHTTP(target string, port string) {
	url := "http://" + target

	if port == "443" {
		url = "https://" + target
	}

	fmt.Println("\n[HTTP] Checking:", url)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("[HTTP] ❌ Request failed")
		fmt.Println("→ Possible causes:")
		fmt.Println(" - web server down")
		fmt.Println(" - TLS/SSL issue")
		fmt.Println(" - firewall or proxy blocking")
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("[HTTP] ✅ Response received")
	fmt.Println("Status:", resp.Status)
}
