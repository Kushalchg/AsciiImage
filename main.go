package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

func main() {
	image := LoadImage()
	resizeImage := ResizeImage(image, 200)

	file, err := os.Create("resize.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)

	}
	defer file.Close()
	png.Encode(file, resizeImage)

	grayImage := ConvGrayScale(resizeImage)
	// for gray image
	grayFile, err := os.Create("gray.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)

	}
	defer file.Close()

	png.Encode(grayFile, grayImage)
	resultStr := MapAscii(grayImage)
	fmt.Printf("MapAscii returned %d lines\n", len(resultStr))
	saveToFile(resultStr, "result.txt")

}
func LoadImage() image.Image {

	file, err := os.Open("me.png")
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

func ResizeImage(img image.Image, width int) image.Image {
	bounds := img.Bounds()
	height := (bounds.Dy() * width) / bounds.Dx()

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Resize the mask image to match the target dimensions
	draw.CatmullRom.Scale(newImage, newImage.Bounds(), img, bounds, draw.Over, nil)

	return newImage
}

func ConvGrayScale(img image.Image) image.Image {
	bound := img.Bounds()

	grayImage := image.NewRGBA(bound)

	for i := bound.Min.X; i < bound.Max.X; i++ {
		for j := bound.Min.Y; j < bound.Max.Y; j++ {
			oldPixel := img.At(i, j)
			color := color.GrayModel.Convert(oldPixel)
			// // fmt.Print(color)
			// r, g, b, _ := oldPixel.RGBA()

			// grayValue := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			// color := color.Gray{uint8(grayValue / 256)}
			grayImage.Set(i, j, color)

		}
	}

	return grayImage
}

func MapAscii(img image.Image) []string {
	// asciiChar := ".`'^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	asciiChar := "$@B%#*+=,.      "
	bound := img.Bounds()
	height, width := bound.Max.Y, bound.Max.X
	result := make([]string, height)

	for y := bound.Min.Y; y < height; y++ {
		line := ""
		for x := bound.Min.X; x < width; x++ {
			pixelValue := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			pixel := pixelValue.Y
			fmt.Print(pixel)
			asciiIndex := int(pixel) * (len(asciiChar) - 1) / 255

			fmt.Print(asciiIndex)
			fmt.Println()
			line += string(asciiChar[asciiIndex])

		}
		result[y] = line

	}
	return result
}

func saveToFile(asciiArt []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range asciiArt {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
