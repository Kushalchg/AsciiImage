package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func main() {
	image := LoadImage()
	resizeImage := ResizeImage(image, 3000)

	file, err := os.Create("logo.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)

	}
	defer file.Close()
	png.Encode(file, resizeImage)

}
func LoadImage() image.Image {

	file, err := os.Open("test.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)
	}
	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		fmt.Printf("error while decoding image %v\n", err)
	}
	return img

}

// func ResizeImage(img image.Image, width int) image.Image {
// 	bound := img.Bounds()
// 	height := (bound.Dy() * width) / bound.Dx()
// 	newImage := image.NewRGBA(image.Rect(0, 0, height, width))
// 	draw.Draw(newImage, newImage.Bounds(), img, img.Bounds().Size(), draw.Src)
// 	return newImage

// }
func ResizeImage(img image.Image, width int) image.Image {
	bounds := img.Bounds()
	height := (bounds.Dy() * width) / bounds.Dx()

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(newImage, newImage.Bounds(), img, bounds, draw.Over, nil)
	return newImage
}
