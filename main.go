package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		help()
		return
	}
	if os.Args[1] == "-l" || os.Args[1] == "--list" {
		getAvaliableScreens()
		return
	}
	if len(os.Args) == 2 {
		screen, err := stringToInt(os.Args[1])
		if err != nil {
			fmt.Println("Invalid screen id")
			return
		}
		result, err := captureScreenshot(screen, "screenshot")
		if err != nil {
			if err.Error() == "GetDIBits failed" {
				fmt.Println("Invalid screen id")
				return
			} else {
				fmt.Printf("Error capturing screenshot: %s\n", err.Error())
				return
			}
		}
		fmt.Println(result)
		return
	}
	if len(os.Args) == 3 {
		screen, err := stringToInt(os.Args[1])
		if err != nil {
			fmt.Println("Invalid screen id")
			return
		}
		result, err := captureScreenshot(screen, os.Args[2])
		if err != nil {
			if err.Error() == "GetDIBits failed" {
				fmt.Println("Invalid screen id")
				return
			} else {
				fmt.Printf("Error capturing screenshot: %s\n", err.Error())
				return
			}
		}
		fmt.Println(result)
		return
	}
	if len(os.Args) == 4 {
		screen, err := stringToInt(os.Args[1])
		if err != nil {
			fmt.Println("Invalid screen id")
			return
		}
		delay, err := stringToInt(os.Args[3])
		if err != nil {
			fmt.Println("Invalid delay")
			return
		}
		countdown(delay)
		time.Sleep(time.Duration(delay) * time.Second)
		result, err := captureScreenshot(screen, os.Args[2])
		if err != nil {
			if err.Error() == "GetDIBits failed" {
				fmt.Println("Invalid screen id")
				return
			} else {
				fmt.Printf("Error capturing screenshot: %s\n", err.Error())
				return
			}
		}
		fmt.Println(result)
		return
	}
}

func help() {
	fmt.Println("Usage: gorecorder [arguments]")
	fmt.Println("Arguments:")
	fmt.Println("-h or --help: Print help")
	fmt.Println("-l or --list: List avaliable screens")
	fmt.Println("(screen id): Capture screenshot from the screen id")
	fmt.Println("(screen id) (output file name): Capture screenshot from the screen id and save it to the output file name")
	fmt.Println("(screen id) (output file name) (delay in seconds): Capture screenshot from the screen id and save it to the output file name after the delay in seconds")
	fmt.Println("Example: gorecorder 0 screenshot 5")
}

func countdown(seconds int) {
	for i := seconds; i > 0; i-- {
		fmt.Printf("Taking in... %d", i)
		time.Sleep(time.Second)
		fmt.Printf("\r")
	}
	fmt.Printf("Say cheese!")
	fmt.Printf("\r")
}

func stringToInt(str string) (int, error) {
	var num int
	_, err := fmt.Sscanf(str, "%d", &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func getAvaliableScreens() {
	n := screenshot.NumActiveDisplays()
	fmt.Printf("Active displays: %d\n", n)
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		fmt.Printf("Id: '%d', Bounds '%v'\n", i, bounds)
	}
}

func captureScreenshot(screen int, output string) (string, error) {
	bounds := screenshot.GetDisplayBounds(screen)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s.jpg", output)
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	jpeg.Encode(file, img, &jpeg.Options{Quality: 80})
	return "Screenshot saved to " + fileName, nil
}
