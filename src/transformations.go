package main

import (
	"image"
	"math/rand"
	"os"
)

func transform() {
	if flagops.inverting {
		invert()
	}

	if flagops.grayscaling {
		gray()
	}
}

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
