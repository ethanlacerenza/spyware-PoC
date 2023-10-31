package main

import (
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	// Retrieve the URL from an environment variable
	serverURL := os.Getenv("SCREENSHOT_SERVER_URL")

	if serverURL == "" {
		panic("The server URL is not specified.")
	}

	for {
		captureAndSendScreenshot(serverURL)
		time.Sleep(3 * time.Second)
	}
}

func captureAndSendScreenshot(serverURL string) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}

	// Now, you can send the image to the server URL.
	// Use `serverURL` to send the image to the remote server.
	// Implement the logic to send to the server as per your preference.
}
