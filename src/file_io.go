package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

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
