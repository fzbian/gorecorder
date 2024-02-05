package main

import (
	"bytes"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
	"golang.design/x/clipboard"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"time"
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
	responseContainer := container.NewVBox(widget.NewLabel(""))

	return container.NewHBox(
		widget.NewButton("Capture and save as", func() {
			dialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
				if writer == nil {
					return
				}
				if err != nil {
					log.Panic(err.Error())
				}
				defer func(writer fyne.URIWriteCloser) {
					err := writer.Close()
					if err != nil {
						log.Panic(err.Error())
					}
				}(writer)
				myWindow.Hide()
				toBytes, err := captureToBytes(actualScreen)
				if err != nil {
					log.Panic(err.Error())
				}
				n, err := writer.Write(toBytes)
				if err != nil {
					log.Panic(err.Error())
				}
				go func() {
					time.Sleep(1 * time.Second)
					myWindow.Show()
				}()
				println(fmt.Sprintf("writed %d bytes", n))
				responseContainer.Objects[0] = widget.NewLabel("Screenshot saved")
				responseContainer.Refresh()
			}, myWindow)
			dialog.SetFileName(fmt.Sprintf("screenshot.%s", defaultFileTypeScreenshot))
			dialog.Show()
		}),
		widget.NewButton("Copy to clipboard", func() {
			myWindow.Hide()
			toBytes, err := captureToBytes(actualScreen)
			if err != nil {
				log.Panic(err.Error())
			}
			err = clipboard.Init()
			if err != nil {
				log.Panic(err.Error())
			}
			clipboard.Write(clipboard.FmtImage, toBytes)
			go func() {
				time.Sleep(1 * time.Second)
				myWindow.Show()
			}()
			responseContainer.Objects[0] = widget.NewLabel("Screenshot copied to clipboard")
			responseContainer.Refresh()
		}),
		responseContainer,
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

func captureToBytes(screen int) ([]byte, error) {
	bounds := screenshot.GetDisplayBounds(screen)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}

	imageBytes, err := imageToBytes(*img)
	if err != nil {
		return nil, err
	}
	return imageBytes, nil
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

func imageToBytes(img image.RGBA) ([]byte, error) {
	var imgBytes []byte
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, &img)
	if err != nil {
		return nil, err
	}
	imgBytes = buffer.Bytes()

	return imgBytes, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}
