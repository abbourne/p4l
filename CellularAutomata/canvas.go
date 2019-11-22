package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"p4l/code.google.com/p/draw2d/draw2d"
)

type Canvas struct {
	gc     *draw2d.ImageGraphicContext
	img    image.Image
	width  int
	height int
}

// Create a new canvas
func CreateNewCanvas(w, h int) Canvas {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(i)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)
	// fill the background
	gc.Clear()
	gc.SetFillColor(image.Black)

	return Canvas{gc, i, w, h}
}

// Create a new color
func MakeColor(r, g, b uint8) color.Color {
	return &color.RGBA{r, g, b, 255}
}

// Move the current point to (x,y)
func (c *Canvas) MoveTo(x, y float64) {
	c.gc.MoveTo(x, y)
}

// Draw a line from the current point to (x,y), and set the current point to (x,y)
func (c *Canvas) LineTo(x, y float64) {
	c.gc.LineTo(x, y)
}

// Set the line color
func (c *Canvas) SetStrokeColor(col color.Color) {
	c.gc.SetStrokeColor(col)
}

// Set the fill color
func (c *Canvas) SetFillColor(col color.Color) {
	c.gc.SetFillColor(col)
}

// Set the line width
func (c *Canvas) SetLineWidth(w float64) {
	c.gc.SetLineWidth(w)
}

// Actually draw the lines you've set up with LineTo
func (c *Canvas) Stroke() {
	c.gc.Stroke()
}

// Fill the area inside the lines you've set up with LineTo
func (c *Canvas) FillStroke() {
	c.gc.FillStroke()
}

// Fill the area inside the lines you've set up with LineTo, but don't
// draw the lines
func (c *Canvas) Fill() {
	c.gc.Fill()
}

// Fill the whole canvas with the fill color
func (c *Canvas) Clear() {
	c.gc.Clear()
}

// Fill the given rectangle with the fill color
func (c *Canvas) ClearRect(x1, y1, x2, y2 int) {
	c.gc.ClearRect(x1, y1, x2, y2)
}

// Save the current canvas to a PNG file
func (c *Canvas) SaveToPNG(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, c.img)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filename)
}

// Return the width of the canvas
func (c *Canvas) Width() int {
	return c.width
}

// Return the height of the canvas
func (c *Canvas) Height() int {
	return c.height
}