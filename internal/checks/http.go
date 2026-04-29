package checks

import (
	"fmt"
	"net/http"
	"time"
)

func CheckHTTP(target string, port string) bool {
	url := "http://" + target

	if port == "443" {
		url = "https://" + target
	}

	fmt.Println("\n[HTTP] Checking:", url)

	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// erlaubt Redirects (wichtig für echte Webseiten)
			return nil
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("[HTTP] ❌ Request build failed")
		fmt.Println("Error:", err)
		return false
	}

	// wirkt wie ein echter Client (wichtig!)
	req.Header.Set("User-Agent", "netcheckup/1.0")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[HTTP] ❌ Request failed")
		fmt.Println("→ Possible causes:")
		fmt.Println(" - web server down")
		fmt.Println(" - TLS/SSL issue")
		fmt.Println(" - firewall or proxy blocking")
		fmt.Println("Error:", err)
		return false
	}
	defer resp.Body.Close()

	fmt.Println("[HTTP] ✅ Response received")
	fmt.Println("Status:", resp.Status)

	// alles unter 500 ist grundsätzlich erreichbar
	if resp.StatusCode < 500 {
		return true
	}

	return false
}
