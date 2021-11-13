package main

import (
	"fmt"
	"math/rand"
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

	img, ftype, nImg := decodeInput()

	if flagops.showdim {
		fmt.Println("Selected image dimensions:", img.Bounds().Max)
	}

	pixels = nImg.Pix

	transform()

	nImg.Pix = pixels

	encodeOutput(nImg, ftype)
}
