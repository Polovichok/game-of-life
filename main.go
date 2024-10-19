package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/png"
	"os"
)

const (
	length         int = 1000
	width          int = 1000
	squareSideSize int = 50
	frameNumber    int = 100
)

func main() {
	answer := createArray()
	fillingArray(&answer)

	for frame := 0; frame < frameNumber; frame++ {

		new_png_file := fmt.Sprintf("e:/prog-projects/go-projects/game-of-life/assets/frame_%d.png", frame)

		_, myImage := createImage()
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
		fillingNewArray(&answer)
	}
}


