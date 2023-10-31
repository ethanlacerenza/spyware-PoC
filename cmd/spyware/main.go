package main

import (
	"os"
	"runtime"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/kbinani/xrandr"
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
	var img *screenshot.Image
	var err error

	// Check the runtime OS
	switch runtime.GOOS {
	case "windows":
		img, err = screenshot.CaptureScreen()
	case "linux":
		outputs, err := xrandr.GetActiveOutputs()
		if err == nil && len(outputs) > 0 {
			img, err = screenshot.CaptureScreen(xrandrScreen(outputs[0]))
		}
	default:
		panic("Unsupported operating system")
	}

	if err != nil {
		panic(err)
	}

	// Now, you can send the image to the server URL.
	// Use `serverURL` to send the image to the remote server.
	// Implement the logic to send to the server as per your preference.
}

// Convert xrandr output to screenshot screen
func xrandrScreen(output xrandr.Output) screenshot.Screen {
	return screenshot.Screen{
		X:      output.Crtc.X,
		Y:      output.Crtc.Y,
		Width:  output.Crtc.Width,
		Height: output.Crtc.Height,
	}
}
