package main

import (
	"image"
	"p4l/canvas"
)

// let's place our drawing functions here.

//AnimateSystem takes a collection of Universe objects along with a canvas width
//parameter and generates a slice of images corresponding to drawing each Universe
//on a canvasWidth x canvasWidth canvas
func AnimateSystem(timePoints []Universe, canvasWidth int, frequency int) []image.Image {
	images := make([]image.Image, 0)
	for i := range timePoints {
		if i%frequency == 0 {
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	}
	return images
}

//DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
//object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels
func DrawToCanvas(u Universe, canvasWidth int) image.Image {
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	for i, b := range u.bodies {
		c.SetFillColor(canvas.MakeColor(b.red, b.green, b.blue))
		cx := (b.position.x / u.width) * float64(canvasWidth)
		cy := (b.position.y / u.width) * float64(canvasWidth)
		r := (b.radius / u.width) * float64(canvasWidth)
		if i == 0 { // For Jupiter, don't scale it

			c.Circle(cx, cy, r)
		} else { //Scale Moons by a factor of 10 to be visible
			r *= 10
			c.Circle(cx, cy, r)
		}
		c.Fill()
	}
	return c.GetImage()
}
