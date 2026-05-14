package checks

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

type PacketLossResult struct {
	Enabled     bool   `json:"enabled"`
	OK          bool   `json:"ok"`
	Sent        int    `json:"sent"`
	Received    int    `json:"received"`
	LossPercent int    `json:"loss_percent"`
	DurationMS  int64  `json:"duration_ms"`
	Error       string `json:"error,omitempty"`
}

func CheckPacketLoss(target string, count int, timeout int, verbose bool) PacketLossResult {
	start := time.Now()

	result := PacketLossResult{
		Enabled: true,
		Sent:    count,
	}

	if verbose {
		fmt.Println()
		fmt.Println("[LOSS] Checking packet loss:", target)
	}

	cmd := exec.Command(
		"ping",
		"-c", strconv.Itoa(count),
		"-W", strconv.Itoa(timeout),
		target,
	)

	output, err := cmd.CombinedOutput()
	result.DurationMS = time.Since(start).Milliseconds()

	if err != nil && len(output) == 0 {
		result.OK = false
		result.Error = err.Error()

		if verbose {
			fmt.Println("[LOSS] ❌ Packet loss check failed")
			fmt.Println("Error:", err)
		}

		return result
	}

	text := string(output)

	packetRegex := regexp.MustCompile(`(\d+) packets transmitted, (\d+) packets received`)
	match := packetRegex.FindStringSubmatch(text)

	if len(match) < 3 {
		result.OK = false
		result.Error = "could not parse ping output"

		if verbose {
			fmt.Println("[LOSS] ❌ Could not parse ping output")
		}

		return result
	}

	sent, _ := strconv.Atoi(match[1])
	received, _ := strconv.Atoi(match[2])

	result.Sent = sent
	result.Received = received

	if sent > 0 {
		result.LossPercent = int(float64(sent-received) / float64(sent) * 100)
	}

	result.OK = result.LossPercent == 0

	if verbose {
		if result.OK {
			fmt.Println("[LOSS] ✅ Packet loss:", result.LossPercent, "%")
		} else {
			fmt.Println("[LOSS] ❌ Packet loss:", result.LossPercent, "%")
		}

		fmt.Printf("[LOSS] Packets: %d sent, %d received\n", result.Sent, result.Received)
	}

	return result
}
