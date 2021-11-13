package main

import "flag"

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
