package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var teal color.Color = color.RGBA{0, 200, 200, 255}

func HorizontalLine(img *image.RGBA, x1, y1, len int, col color.Color) {

	for x := x1; x < (x1 + len); x++ {
		img.Set(x, y1, col)
	}
}
func main() {
	file, err := os.Create("someimage.png")

	if err != nil {
		fmt.Errorf("%s", err)
	}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)
	// или draw.Draw(img, img.Bounds(), image.Transparent, image.ZP, draw.Src)
	var red color.Color = color.RGBA{200, 30, 30, 255}
	HorizontalLine(img, 50, 50, 100, red)

	png.Encode(file, img)
}
