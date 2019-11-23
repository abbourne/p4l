package main

import (
	"fmt"
	"p4l/vec"
)

// Manual tests of the vec package
func main() {
	p1 := vec.Point{3, 4, 0}
	p1Len := p1.Length()
	fmt.Println("length of", p1, "is", p1Len)
}
