package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

var actualScreen int
var defaultQualityScreenshot = 80

var myApp = app.New()
var myWindow = myApp.NewWindow("ScreenGo")

func main() {
	myWindow.Resize(fyne.NewSize(520, 320))
	myWindow.SetFixedSize(true)
	myWindow.SetContent(container.NewVBox(
		selectWindowContainer(),
		widget.NewSeparator(),
		selectQualityContainer(),
		widget.NewSeparator(),
		selectFileTypeContainer(),
		widget.NewSeparator(),
		captureWindowContainer(),
	))
	myWindow.ShowAndRun()
}

func selectWindowContainer() *fyne.Container {
	screensStr := getAvaliableScreens()
	windowSelect := widget.NewSelect(screensStr, func(value string) {
		for i, screen := range screensStr {
			if screen == value {
				actualScreen = i
				break
			}
		}
	})
	windowSelect.Selected = screensStr[0]
	return container.NewVBox(
		widget.NewLabel("Select a screen"),
		windowSelect,
	)
}

var defaultFileTypeScreenshot = "png"

func selectFileTypeContainer() *fyne.Container {
	fileTypes := widget.NewRadioGroup([]string{"png", "jpg"}, func(value string) {
		switch value {
		case "png":
			defaultFileTypeScreenshot = "png"
		case "jpg":
			defaultFileTypeScreenshot = "jpg"
		}
	})
	fileTypes.Selected = "png"
	fileTypes.Horizontal = true
	return container.NewVBox(
		widget.NewLabel("Select the filetype"),
		fileTypes)
}

func captureWindowContainer() *fyne.Container {
	output := widget.NewEntry()
	output.SetPlaceHolder("Output file name (default: screenshot.jpg)")
	responseContainer := container.NewVBox(widget.NewLabel(""))
	return container.NewVBox(
		widget.NewLabel("Output file name"),
		output,
		widget.NewButton("Capture", func() {
			msg, err := captureScreenshot(actualScreen, output.Text)
			if err != nil {
				responseContainer.Objects[0] = widget.NewLabel(err.Error())
				responseContainer.Refresh()
			}
			responseContainer.Objects[0] = widget.NewLabel(msg)
			responseContainer.Refresh()
		}), responseContainer,
	)
}

func selectQualityContainer() *fyne.Container {
	quality := widget.NewRadioGroup([]string{"Low", "Medium", "High"}, func(value string) {
		switch value {
		case "Low":
			defaultQualityScreenshot = 10
		case "Medium":
			defaultQualityScreenshot = 50
		case "High":
			defaultQualityScreenshot = 80
		}
	})
	quality.Selected = "High"
	quality.Horizontal = true
	return container.NewVBox(
		widget.NewLabel("Select a quality"),
		quality,
	)
}

func getAvaliableScreens() []string {
	n := screenshot.NumActiveDisplays()
	screensStr := make([]string, n)
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		infoScreen := fmt.Sprintf("Id: %d, Bounds: %v", i, bounds)
		screensStr[i] = infoScreen
	}

	return screensStr
}

func captureScreenshot(screen int, fileName string) (string, error) {
	bounds := screenshot.GetDisplayBounds(screen)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return "", err
	}

	fileCreatorResponse, err := createFile(fileName, img, defaultFileTypeScreenshot)
	if err != nil {
		return "", nil
	}

	return fileCreatorResponse, nil
}

func createFile(fileName string, img *image.RGBA, fileType string) (string, error) {
	if fileName == "" {
		fileName = "screenshot"
	}

	switch fileType {
	case "jpg":
		attempt := 1
		baseFileName := fileName
		for fileExists(fileName + ".jpg") {
			fileName = fmt.Sprintf("%s (%d)", baseFileName, attempt)
			attempt++
		}

		file, err := os.Create(fileName + ".jpg")
		if err != nil {
			return "", err
		}
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				panic(err.Error())
			}
		}(file)

		err = jpeg.Encode(file, img, &jpeg.Options{Quality: defaultQualityScreenshot})
		if err != nil {
			return "", err
		}
	case "png":
		attempt := 1
		baseFileName := fileName
		for fileExists(fileName + ".png") {
			fileName = fmt.Sprintf("%s (%d)", baseFileName, attempt)
			attempt++
		}

		file, err := os.Create(fileName + ".png")
		if err != nil {
			return "", err
		}
		defer func(file *os.File) {
			err = file.Close()
			if err != nil {
				panic(err.Error())
			}
		}(file)

		err = png.Encode(file, img)
		if err != nil {
			return "", err
		}
	}

	return "Screenshot saved to " + fileName + "." + fileType, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}
