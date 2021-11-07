package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
)

type Pixel struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func gray(arr []uint8, i int) Pixel {
	avg := uint8(float32(arr[i])*0.3) + uint8(float32(arr[i+1])*0.59) + uint8(float32(arr[i+2])*0.11)
	return Pixel{avg, avg, avg, 255}
}

func errorCheck(err error, a ...interface{}) {
	if err != nil {
		fmt.Println("error: ", err, a)
	}
}

func main() {
	fImg, err := os.Open("./images/4k_mountains.jpg")
	errorCheck(err, "opening")
	defer fImg.Close()
	fImg.Seek(0, 0)

	img, err := jpeg.Decode(fImg)
	errorCheck(err, "decode")
	fImg.Seek(0, 0)

	nImg := image.NewRGBA(img.Bounds())
	// Copy the image pixels into nImg
	draw.Draw(nImg, img.Bounds(), img, img.Bounds().Min, draw.Src)

	result := []uint8{}
	for i := 0; i < len(nImg.Pix); i += 4 {
		result = append(result, gray(nImg.Pix, i).r, gray(nImg.Pix, i).g, gray(nImg.Pix, i).b, gray(nImg.Pix, i).a)
	}
	nImg.Pix = result

	oImg, err := os.Create("out.jpg")
	errorCheck(err)
	jpeg.Encode(oImg, nImg, nil)
}
