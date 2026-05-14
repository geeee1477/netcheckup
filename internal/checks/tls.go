package checks

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type TLSResult struct {
	Enabled    bool   `json:"enabled"`
	OK         bool   `json:"ok"`
	Issuer     string `json:"issuer,omitempty"`
	Subject    string `json:"subject,omitempty"`
	ExpiresAt  string `json:"expires_at,omitempty"`
	DaysLeft   int    `json:"days_left,omitempty"`
	Error      string `json:"error,omitempty"`
	DurationMS int64  `json:"duration_ms,omitempty"`
}

func CheckTLS(target string, port string, timeoutSeconds int, verbose bool) TLSResult {
	start := time.Now()

	result := TLSResult{
		Enabled: true,
	}

	address := target + ":" + port

	if verbose {
		fmt.Println("\n[TLS] Checking:", address)
	}

	conn, err := tls.DialWithDialer(
		&net.Dialer{
			Timeout: time.Duration(timeoutSeconds) * time.Second,
		},
		"tcp",
		address,
		&tls.Config{
			ServerName: target,
		},
	)

	if err != nil {
		result.OK = false
		result.Error = err.Error()
		result.DurationMS = time.Since(start).Milliseconds()

		if verbose {
			fmt.Println("[TLS] ❌ Certificate check failed")
			fmt.Println("Error:", err)
		}

		return result
	}

	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		result.OK = false
		result.Error = "no certificate returned"
		result.DurationMS = time.Since(start).Milliseconds()
		return result
	}

	cert := certs[0]
	daysLeft := int(time.Until(cert.NotAfter).Hours() / 24)

	result.OK = time.Now().Before(cert.NotAfter)
	result.Issuer = cert.Issuer.CommonName
	result.Subject = cert.Subject.CommonName
	result.ExpiresAt = cert.NotAfter.Format(time.RFC3339)
	result.DaysLeft = daysLeft
	result.DurationMS = time.Since(start).Milliseconds()

	if verbose {
		fmt.Println("[TLS] ✅ Certificate valid")
		fmt.Println("[TLS] Issuer:", result.Issuer)
		fmt.Println("[TLS] Subject:", result.Subject)
		fmt.Println("[TLS] Expires in:", result.DaysLeft, "days")
	}

	return result
}
