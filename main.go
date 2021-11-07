package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

func errorCheck(err error, a ...interface{}) {
	if err != nil {
		fmt.Println("error: ", err, a)
	}
}

func main() {
	fImg, err := os.Open("./images/bee.jpg")
	errorCheck(err, "opening")
	defer fImg.Close()
	fImg.Seek(0, 0)

	img, err := jpeg.Decode(fImg)
	errorCheck(err, "decode")
	fImg.Seek(0, 0)

	nImg := image.NewRGBA(img.Bounds())
	draw.Draw(nImg, img.Bounds(), img, img.Bounds().Min, draw.Src)

	oImg, err := os.Create("out.jpg")
	errorCheck(err)
	jpeg.Encode(oImg, nImg, nil)
}
