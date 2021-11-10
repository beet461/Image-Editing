package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
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

func gray() {
	for i := 0; i < len(pixels); i += 4 {
		avg := uint8(float32(pixels[i])*0.3) + uint8(float32(pixels[i+1])*0.59) + uint8(float32(pixels[i+2])*0.11)
		pixels[i] = avg
		pixels[i+1] = avg
		pixels[i+2] = avg
		pixels[i+3] = 255
	}
}

func invert() {
	for i := 0; i < len(pixels); i += 4 {
		pixels[i] = 255 - pixels[i]
		pixels[i+1] = 255 - pixels[i+1]
		pixels[i+2] = 255 - pixels[i+2]
	}
}

func randomnoisegen() {
	for i := 0; i < flagops.outputimagedim.Height*flagops.outputimagedim.Width; i++ {
		pixels = append(pixels, uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255)
	}
	nImg := image.NewRGBA(image.Rect(0, 0, flagops.outputimagedim.Width, flagops.outputimagedim.Height))
	nImg.Pix = pixels
	encodeOutput(nImg, "jpeg")
	os.Exit(0)
}

func errorCheck(err error, a ...interface{}) {
	if err != nil {
		fmt.Println(a, "error: ", err)
	}
}

func encodeOutput(nImg image.Image, ftype string) {
	oImg, err := os.Create(flagops.outputfile)
	errorCheck(err)
	switch ftype {
	case "jpeg":
		jpeg.Encode(oImg, nImg, nil)
	case "png":
		png.Encode(oImg, nImg)
	}
}

func decodeInput() (image.Image, string) {
	fImg, err := os.Open(flagops.inputfile)
	errorCheck(err, "opening")
	defer fImg.Close()
	fImg.Seek(0, 0)

	img, ftype, err := image.Decode(fImg)
	errorCheck(err, "decode")
	fImg.Seek(0, 0)

	switch ftype {
	case "jpeg":
		img, err = jpeg.Decode(fImg)
	case "png":
		img, err = png.Decode(fImg)
	}
	errorCheck(err, "decode")
	fImg.Seek(0, 0)

	return img, ftype
}

func set_flags() {
	flag.StringVar(&flagops.inputfile, "input", "", "Path to the input file")
	flag.BoolVar(&flagops.showdim, "showdim", false, "Print the dimensions of the input image")
	flag.StringVar(&flagops.outputfile, "output", "./out.jpg", "Path to the output image")
	flag.BoolVar(&flagops.grayscaling, "grayscale", false, "Make an image grayscaled")
	flag.BoolVar(&flagops.inverting, "invertcolours", false, "Invert the image colours")
	flag.BoolVar(&flagops.randomnoise, "randomnoise", false, "Generate an image where each pixel is randomly generated. Set dimensions with -outputheight and -outputwidth")
	flag.IntVar(&flagops.outputimagedim.Height, "outputheight", 720, "Height of the output image in pixels")
	flag.IntVar(&flagops.outputimagedim.Width, "outputwidth", 1260, "Width of the output image in pixels") // random noise generation dimensions set it
	flag.Parse()
}

func main() {
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
