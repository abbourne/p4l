package main

import (
	"fmt"
	"gogif"
	"image"
	"image/gif"
	"os"
)

// process() takes a slice of images and uses them to generate an animated GIF
// with the name "filename.out.gif" where filename is an input parameter.
func process(imglist []image.Image, filename string) {

	// get ready to write images to files
	w, err := os.Create(filename + ".out.gif")

	if err != nil {
		fmt.Println("Sorry: couldn't create the file!")
		os.Exit(1)
	}

	defer w.Close()
	var g gif.GIF
	g.Delay = make([]int, len(imglist))
	g.Image = make([]*image.Paletted, len(imglist))
	g.LoopCount = 10

	for i := range imglist {
		g.Image[i] = ImageToPaletted(imglist[i])
		g.Delay[i] = 1
	}

	gif.EncodeAll(w, &g)
}

// ImageToPaletted converts an image to an image.Paletted with 256 colors.
// It is used by a subroutine by process() to generate an animated GIF.
func ImageToPaletted(img image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)
	if !ok {
		b := img.Bounds()
		pm = image.NewPaletted(b, nil)
		q := &gogif.MedianCutQuantizer{NumColor: 256}
		q.Quantize(pm, b, img, image.ZP)
	}
	return pm
}
