package main

import (
	"fmt"

	"github.com/llgcode/draw2d/draw2dkit"
)

// Add your drawing functions here! :)

func main_old() {

	red := MakeColor(255, 0, 0)
	//blue := MakeColor(0, 0, 255)
	//green := MakeColor(0, 255, 0)
	//purple := MakeColor(128, 0, 128)
	black := MakeColor(0, 0, 0)
	//white := MakeColor(255, 255, 255)

	fmt.Println("Drawing lab!")
	pic := CreateNewCanvas(1000, 300)
	pic.SetStrokeColor(red)
	pic.MoveTo(100, 100)
	pic.LineTo(100, 200)
	pic.LineTo(900, 200)
	pic.LineTo(900, 100)
	pic.LineTo(100, 100)
	pic.Stroke()

	//func Rectangle(path draw2d.PathBuilder, x1, y1, x2, y2 float64)
	pic.SetStrokeColor(black)
	pic.SetLineWidth(30)
	draw2dkit.Rectangle(pic.gc, 75, 75, 250, 250)
	pic.Stroke()
	pic.SaveToPNG("MyRedRectangle.png")

}
