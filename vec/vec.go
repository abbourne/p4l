// Package vec contains basic vector math operations.
package vec

import "math"

// Point is a single point in 3D space
type Point struct {
	X float64
	Y float64
	Z float64
}

// Vec is a 3 dimensional vector.
// It's represented as a point relative to the origin
// But it's implementation should be considered as private
type Vec struct {
	Point
}

// Origin is the point representing the zeo origin in 3D space. Treat it as a const!
var Origin = Point{0, 0, 0}

// Distance returns the distance between 2 points in space
func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p.X, 2) + math.Pow(p2.Y-p.Y, 2) + math.Pow(p2.Z-p.Z, 2))
}

// Length returns the magnitude of a Vec, or the distance of a Point from the Origin
func (p Point) Length() float64 {
	return Origin.Distance(p)
}

// Minus subtracts a point from the receiver point and returns a new Point This is also vector subtraction
func (p Point) Minus(p2 Point) Point {
	x := p.X - p2.X
	y := p.Y - p2.Y
	z := p.Z - p2.Z
	return Point{x, y, z}
}

// Plus adds a point to the receiver point and returns a new Point This is also vector addition
func (p Point) Plus(p2 Point) Point {
	x := p.X + p2.X
	y := p.Y + p2.Y
	z := p.Z + p2.Z
	return Point{x, y, z}
}

// Al gets alpha, the angle to the x axis
func (v Vec) Al() float64 {
	return math.Acos(v.X / v.Length())
}

// Bt gets beta, the angle to the y axis
func (v Vec) Bt() float64 {
	return math.Acos(v.Y / v.Length())
}

// Ga gets gamma, the angle to the z axis
func (v Vec) Ga() float64 {
	return math.Acos(v.Z / v.Length())
}

// I is the unit vector on the x axis. Treat it as a const!
var I = Vec{Point{1, 0, 0}}

// J is the unit vector on the y axis. Treat it as a const!
var J = Vec{Point{0, 1, 0}}

// K is the unit vector on the z axis. Treat it as a const!
var K = Vec{Point{0, 0, 1}}

// Newc creates new Vec vector from a set of x, y, z coordinates
func Newc(x, y, z float64) Vec {
	return Vec{Point{x, y, z}}
}

// Newv creates new Vec vector from an existing Vec
func Newv(v Vec) Vec {
	v2 := v
	return v2
}

// Newd creates new Vec vector from a magnitude and the direction angles
// m is the magnitude of the new vector, al is alpha the angle to the x axis,
// b is beta the angle to the y axis, and gs is gamma the angle to the z axis
func Newd(m, al, bt, ga float64) Vec {
	// Direction cosines are as follows
	//  cos(al) = m.x/m, cos(bt) = m.y/m cos(ga) = m.z/m
	x := math.Cos(al) * m
	y := math.Cos(bt) * m
	z := math.Cos(ga) * m
	return Vec{Point{x, y, z}}
}

// Newp creates a new Vec vector from 2 points. The Vec represents the distance and angle from p1 to p1
func Newp(p1, p2 Point) Vec {
	return Vec{p2.Minus(p1)}
}
