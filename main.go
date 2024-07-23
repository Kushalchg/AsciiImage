package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	// image := LoadImage()
	// image := CreateSampleImage()
	// file, err := os.Create("logo.png")
	// if err != nil {
	// 	fmt.Printf("error while opening file %v\n", err)

	// }
	// png.Encode(file, image)

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

// }
