package main

import (
	"fmt"
	"image"
	"image/draw"
	"math/rand"
	"os"
	"time"
)

type Dim struct {
	Height int
	Width  int
}

type FlagOptions struct {
	inputfile      string
	outputfile     string
	showdim        bool
	grayscaling    bool
	inverting      bool
	randomnoise    bool
	outputimagedim Dim
}

var pixels = []uint8{}
var flagops FlagOptions

func errorCheck(err error, a ...interface{}) {
	if err != nil {
		fmt.Println(a, "error: ", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// Take the input flags from cli
	set_flags()

	if flagops.randomnoise {
		randomnoisegen()
	}

	if flagops.inputfile == "" {
		fmt.Println("WARNING: Input image path is unselected, to select file add -input='yourfile.jpg' at the end of the command")
		os.Exit(0)
	}

	img, ftype := decodeInput()

	if flagops.showdim {
		fmt.Println("Selected image dimensions:", img.Bounds().Max)
	}

	nImg := image.NewRGBA(img.Bounds())
	// Copy the image pixels into nImg
	draw.Draw(nImg, img.Bounds(), img, img.Bounds().Min, draw.Src)

	pixels = nImg.Pix

	if flagops.inverting {
		invert()
	}

	if flagops.grayscaling {
		gray()
	}

	nImg.Pix = pixels

	encodeOutput(nImg, ftype)
}
