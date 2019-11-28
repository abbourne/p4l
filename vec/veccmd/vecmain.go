package main

import (
	"fmt"
	"log"
	"math"
	"p4l/vec"
)

// Manual tests of the vec package
func main() {
	log.SetFlags(log.Lshortfile)
	p1 := vec.Newc(3, 4, 0)
	p1Len := p1.Length()
	fmt.Println("length of", p1, "is", p1Len)

	v1 := vec.Newp(vec.Origin, p1)
	v1Len := v1.Length()
	fmt.Println("length of", v1, "is", v1Len)

	fmt.Println("pi radians is", vec.ToDegrees(math.Pi), "degrees")
	fmt.Println("180 degrees is", vec.ToRadians(180), "radians")

	vec.I.Info("Basis I")
	vec.I.Mul(2).Info("I*2")
	vec.J.Info("Basis J")
	vec.J.Mul(3).Info("J*3")
	vec.K.Info("Basis K")

	vUnit := vec.Newd(1, vec.ToRadians(45), vec.ToRadians(45))
	vUnit.Info("Unit @45deg")
	vUnit = vec.Newd(1, vec.ToRadians(45), vec.ToRadians(90))
	vUnit.Info("Unit in xy plane")
	vUnit = vec.Newc(1, 1, 0)
	vUnit.Info("1:1 in xy plane")
	vUnit.Mul(3).Info("1:1 Multipled by 3")

	aVec := vec.Newc(.5, 1.5, 0)
	c := vUnit.Dot(aVec)
	log.Printf("Dot product is: %G", c)

	vUnit.Cross(aVec).Info("Cross Product")
}
