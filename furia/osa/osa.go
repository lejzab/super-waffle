package osa

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func DrawBackground(img *image.RGBA, c color.Color) {
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, c)
		}
	}
}

func DrawRect(img *image.RGBA, x, y, width, height int, c color.Color) {

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {

			// Get current color of the pixel
			original := img.At(x+i, y+j).(color.RGBA)

			// Calculate the average color
			avgColor := color.RGBA{
				uint8((uint16(original.R) + uint16(c.(color.RGBA).R)) / 2),
				uint8((uint16(original.G) + uint16(c.(color.RGBA).G)) / 2),
				uint8((uint16(original.B) + uint16(c.(color.RGBA).B)) / 2),
				255,
			}

			// Set the new averaged color
			img.Set(x+i, y+j, avgColor)
		}
	}
}

func SaveImage(img image.Image, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}
}

func RandomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
func RandomRectDimensions(maxWidth, maxHeight int) (int, int) {
	return rand.Intn(maxWidth), rand.Intn(maxHeight)
}
func DrawRandomRectangles(img *image.RGBA, count int) {
	for i := 0; i < count; i++ {
		c := RandomColor()
		rectWidth, rectHeight := RandomRectDimensions(150, 50)
		x := rand.Intn(img.Bounds().Dx() - rectWidth)
		y := rand.Intn(img.Bounds().Dy() - rectHeight)
		DrawRect(img, x, y, rectWidth, rectHeight, c)
	}
}
