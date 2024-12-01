package main

import (
	"furia/osa"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"strconv"
	"time"
)

const (
	width  = 800
	height = 600
)

func main() {
	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill background with black color
	backgroundColor := color.RGBA{A: 255}
	osa.DrawBackground(img, backgroundColor)

	// Create and run the fyne application
	myApp := app.New()
	myWindow := myApp.NewWindow("OSA generator")

	// Convert image.RGBA to *canvas.Image
	imageCanvas := canvas.NewImageFromImage(img)
	imageCanvas.FillMode = canvas.ImageFillContain

	// Create label for status bar
	statusLabel := widget.NewLabel("Iteration: 0")

	content := container.NewBorder(nil, statusLabel, nil, nil, imageCanvas)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(width, height))

	go func() {
		for i := 0; i < 100; i++ {
			osa.DrawRandomRectangles(img, 10)
			imageCanvas.Refresh()
			statusLabel.SetText("Iteration: " + strconv.Itoa(i+1))
			time.Sleep(time.Millisecond * 100)
		}
	}()

	myWindow.ShowAndRun()
}
