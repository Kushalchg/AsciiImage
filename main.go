package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	image := LoadImage()
	// greater the image size more clear photo will produce
	resizeImage := ResizeImage(image, 200)

	file, err := os.Create("output/resize.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)
	}
	defer file.Close()
	png.Encode(file, resizeImage)
	grayImage := ConvGrayScale(resizeImage)
	// for gray image
	grayFile, err := os.Create("output/gray.png")
	if err != nil {
		fmt.Printf("error while opening file %v\n", err)

	}
	defer file.Close()
	png.Encode(grayFile, grayImage)
	resultStr := MapAscii(grayImage)
	saveToFile(resultStr, "output/result.txt")

	AsciiToHTML(resultStr)
	AsciiToImage(resultStr)

}
func LoadImage() image.Image {
	file, err := os.Open("output/me.png")
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
	asciiChar := "$@B%#*+=,.       "
	// runes := []rune(asciiChar)

	// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	// 	runes[i], runes[j] = runes[j], runes[i]
	// }

	bound := img.Bounds()
	height, width := bound.Max.Y, bound.Max.X
	result := make([]string, height)

	for y := bound.Min.Y; y < height; y++ {
		line := ""
		for x := bound.Min.X; x < width; x++ {
			pixelValue := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			pixel := pixelValue.Y
			asciiIndex := int(pixel) * (len(asciiChar) - 1) / 255
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

func AsciiToHTML(ascii []string) {
	HtmlFile, err := os.Create("output/format.html")
	if err != nil {
		fmt.Println("error while creatig html file")
	}

	for lin, lines := range ascii {
		htmlString := `<!DOCTYPE html>
		<html lang="en"><head>
   	 	<meta charset="UTF-8">
    	<meta name="viewport" content="width=device-width, initial-scale=0.8">
    	<title>AsciiImage</title>
		</head>
		<body>
			<code>
		 		<span class="ascii" style="color: black;
		  		background: white;
		  		display:inline-block;
		  		white-space:pre;
		  		letter-spacing:0;
		  		line-height:0.9;
		  		font-family:'Consolas','BitstreamVeraSansMono','CourierNew',Courier,monospace;
		  		font-size:10px;
		  		border-width:1px;
		  		border-style:solid;
		  		border-color:lightgray;">`
		if lin == 0 {
			_, err := HtmlFile.WriteString(htmlString)
			if err != nil {
				fmt.Println("error while start writing into html file")
			}
		}

		for _, char := range lines {
			_, err := HtmlFile.WriteString(fmt.Sprintf("<span>%v</span>", string(char)))
			if err != nil {
				fmt.Println("error while writing into html file")
			}
		}
		_, err := HtmlFile.WriteString("<br>")
		if err != nil {
			fmt.Println("error while writing into html file")
		}
		if lin == len(ascii)-1 {
			_, err := HtmlFile.WriteString("</code></body></html>")
			if err != nil {
				fmt.Println("error while end writing into html file")
			}

		}
	}
}

// func SampleImage() {
// 	SampleImage := image.NewRGBA(image.Rect(0, 0, 2, 2))
// 	SampleImage.Set(0, 0, color.RGBA{255, 0, 0, 255})
// 	SampleImage.Set(1, 0, color.RGBA{0, 0, 255, 255})
// 	SampleImage.Set(0, 1, color.RGBA{0, 255, 0, 255})
// 	SampleImage.Set(1, 1, color.RGBA{0, 0, 0, 255})

// 	fmt.Printf("the Sample Image is  %v\n", SampleImage)

// 	file, err := os.Create("output/sample.png")
// 	if err != nil {
// 		fmt.Print("error while creating sample file")
// 	}
// defer file.close()
// 	png.Encode(file, SampleImage)

// }

func AsciiToImage(strArray []string) {
	fmt.Printf("the size is %v", len(strArray))
	// Create a larger image to fit the text
	fontImage := image.NewRGBA(image.Rect(0, 0, 1400, len(strArray)*11))
	// backgroundColor := color.RGBA{0, 0, 255, 255}
	draw.Draw(fontImage, fontImage.Bounds(), image.White, image.Point{}, draw.Src)
	// draw.Draw(fontImage, fontImage.Bounds(), image.NewUniform(backgroundColor), image.Point{}, draw.Src)

	// Create the font face
	drawconf := &font.Drawer{
		Dst:  fontImage,
		Src:  image.Black,
		Face: basicfont.Face7x13,
	}
	// Draw the string
	for i, line := range strArray {
		drawconf.Dot = fixed.Point26_6{
			X: fixed.Int26_6(10 * 64),          // 10 pixels from left
			Y: fixed.Int26_6((20 + i*11) * 64), // Start at 20 pixels from top, then move down by lineHeight for each line
		}

		drawconf.DrawString(line)

	}

	// Create the output file
	file, err := os.Create("output/custom.png")
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()

	// Encode and save the image
	err = png.Encode(file, fontImage)
	if err != nil {
		fmt.Println("Error encoding image:", err)
	}
}
