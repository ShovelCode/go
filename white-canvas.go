package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Create 100x100 white canvas
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	white := color.RGBA{255, 255, 255, 255}

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			img.Set(x, y, white)
		}
	}

	// Save to file
	file, _ := os.Create("out.png")
	defer file.Close()
	png.Encode(file, img)
}
