package main

import (
	"image"
	"image/draw"
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

func decodeInput() (image.Image, string, *image.RGBA) {
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

	nImg := image.NewRGBA(img.Bounds())
	// Copy the image pixels into nImg
	draw.Draw(nImg, img.Bounds(), img, img.Bounds().Min, draw.Src)

	return img, ftype, nImg
}
