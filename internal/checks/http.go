package checks

import (
	"fmt"
	"net/http"
	"time"
)

func CheckHTTP(target string, port string, timeoutSeconds int, retries int, verbose bool) (bool, int64) {
	start := time.Now()

	url := "https://" + target
	if port != "443" {
		url = "http://" + target + ":" + port
	}

	timeout := time.Duration(timeoutSeconds) * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	if verbose {
		fmt.Println("\n[HTTP] Checking:", url)
	}

	for attempt := 1; attempt <= retries; attempt++ {
		resp, err := client.Get(url)
		if err == nil {
			defer resp.Body.Close()

			if verbose {
				fmt.Println("[HTTP] ✅ Response received")
				fmt.Println("Status:", resp.Status)
			}

			return resp.StatusCode < 500, time.Since(start).Milliseconds()
		}

		if verbose {
			fmt.Printf("[HTTP] Attempt %d/%d failed\n", attempt, retries)
			fmt.Println("Error:", err)
		}
	}

	if verbose {
		fmt.Println("[HTTP] ❌ Request failed")
		fmt.Println("→ Possible causes:")
		fmt.Println(" - web server down")
		fmt.Println(" - TLS/SSL issue")
		fmt.Println(" - firewall or proxy blocking")
	}

	return false, time.Since(start).Milliseconds()
}
