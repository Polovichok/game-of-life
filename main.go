package main

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
)

const (
	length         int = 500
	width          int = 500
	squareSideSize int = 25
)

func main() {
	answer := createArray()
	fillingArray(&answer)

	new_png_file, myImage := createField()
	drawSquares(&answer, myImage)

	dc := gg.NewContext(length, width)
	dc.DrawImage(myImage, 0, 0)

	DrawLines(myImage, dc)

	myfile, err := os.Create(new_png_file)
	if err != nil {
		panic(err)
	}

	defer myfile.Close()

	if err := png.Encode(myfile, dc.Image()); err != nil {
		panic(err)
	}
}

func createArray() [width][length]int {
	var array = [width][length]int{}
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			array[i][j] = 0
		}
	}

	return array
}

func fillingArray(array *[width][length]int) {
	for i := 0; i < width; i++ {
		for j := 0; j < length; j++ {
			array[i][j] = rand.Intn(2)
		}
	}
}

func createField() (string, *image.RGBA) {
	new_png_file := "e:/prog-projects/go-projects/game-of-life/rectangle.png"
	myImage := image.NewRGBA(image.Rect(0, 0, length, width))
	myGrey := color.RGBA{160, 160, 160, 160}
	draw.Draw(myImage, myImage.Bounds(), &image.Uniform{myGrey}, image.Point{}, draw.Src)
	return new_png_file, myImage
}

func drawSquares(answer *[width][length]int, myImage *image.RGBA) {
	myWhite := color.RGBA{255, 255, 255, 255}
	myBlack := color.RGBA{0, 0, 0, 255}
	x_0 := 0
	y_0 := 0

	for x_0 < length/squareSideSize && y_0 < width/squareSideSize {
		for i := 0; i < width; i++ {
			for j := 0; j < length; j++ {
				if answer[i][j] == 0 {
					x_0 = squareSideSize * j
					y_0 = squareSideSize * i
					whiteRect := image.Rect(x_0, y_0, x_0+squareSideSize, y_0+squareSideSize)
					draw.Draw(myImage, whiteRect, &image.Uniform{myWhite}, image.Point{}, draw.Src)
				} else {
					x_0 := squareSideSize * j
					y_0 := squareSideSize * i
					blackRect := image.Rect(x_0, y_0, x_0+squareSideSize, y_0+squareSideSize)
					draw.Draw(myImage, blackRect, &image.Uniform{myBlack}, image.Point{}, draw.Src)
				}
			}
		}
	}
}

func DrawLines(myImage *image.RGBA, dc *gg.Context){
	for i := 0; i <= length; i += squareSideSize {

		dc.DrawLine(float64(i), float64(0), float64(i), float64(width))
	}
	for i := 0; i <= width; i += squareSideSize {
		dc.DrawLine(float64(0), float64(i), float64(length), float64(i))
	}
	dc.SetRGB(0,0,0)
	dc.SetLineWidth(1)
	dc.Stroke()
}
