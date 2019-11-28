// Package vec contains basic vector math operations.
package vec

import (
	"log"
	"math"
)

// Vec is a 3 dimensional vector.
// It's represented as a point relative to the origin
// So we also use Vec as if they were points.
// But it's implementation should be considered as private
type Vec struct {
	X float64
	Y float64
	Z float64
}

// Origin is the point representing the zeo origin in 3D space. Treat it as a const!
var Origin = Vec{0, 0, 0}

// Distance returns the distance between 2 points in space
func (v Vec) Distance(v2 Vec) float64 {
	return math.Sqrt(math.Pow(v2.X-v.X, 2) + math.Pow(v2.Y-v.Y, 2) + math.Pow(v2.Z-v.Z, 2))
}

// Length returns the magnitude of a Vec, or the distance of a Vec from the Origin
func (v Vec) Length() float64 {
	return Origin.Distance(v)
}

// AsSpherical returns the spherical coordinates of a Point returning length, theta, phi
// where theta is the angle from the the x axis in the xy plane and
// phi is the angle from the xy plane toward the z axis.
func (v Vec) AsSpherical() (m float64, theta float64, phi float64) {
	m = v.Length()
	phi = math.Acos(v.Z / m)
	theta = math.Asin(v.Y / (m * math.Sin(phi)))
	return
}

// Minus subtracts a Vec from the receiver Vec and returns a new Vec This is also vector subtraction
func (v Vec) Minus(v2 Vec) Vec {
	x := v.X - v2.X
	y := v.Y - v2.Y
	z := v.Z - v2.Z
	return Vec{x, y, z}
}

// Plus adds a Vec to the receiver Vec and returns a new Vec This is also vector addition
func (v Vec) Plus(v2 Vec) Vec {
	x := v.X + v2.X
	y := v.Y + v2.Y
	z := v.Z + v2.Z
	return Vec{x, y, z}
}

// ToDegrees converts from degrees to radians
// 2*pi radians = 360 degrees
func ToDegrees(rad float64) float64 {
	return (rad * 180 / math.Pi)
}

// ToRadians converts from radians to degrees
// 2*pi radians = 360 degrees
func ToRadians(deg float64) float64 {
	return (deg * math.Pi / 180)
}

// Info prints a bunch of info for debugging purposes
func (v Vec) Info(str string) {
	m, th, ph := v.AsSpherical()
	log.Printf("%s:: Vector %v length: %7.3G, theta: %7.3G, phi: %7.3G degrees", str, v, m,
		ToDegrees(th), ToDegrees(ph))
}

// I is the unit vector on the x axis. Treat it as a const!
var I = Vec{1, 0, 0}

// J is the unit vector on the y axis. Treat it as a const!
var J = Vec{0, 1, 0}

// K is the unit vector on the z axis. Treat it as a const!
var K = Vec{0, 0, 1}

// Newc creates new Vec vector from a set of x, y, z coordinates
func Newc(x, y, z float64) Vec {
	return Vec{x, y, z}
}

// Newv creates new Vec vector from an existing Vec
func Newv(v Vec) Vec {
	v2 := v
	return v2
}

// Newd creates new Vec vector from a magnitude and the direction angles
// m is the magnitude of the new vector, th is theta the angle from the x axis in the xy plane
// ph is phi the angle fron the xy plane toward the z axis.
func Newd(m, th, ph float64) Vec {
	// Conversion from cylndrical to cartesian coordinates are as follows:
	// x = m*sin(ph)*cos(th) y = m*sin(ph)*sin(th) z= m*cos(ph)
	x := math.Sin(ph) * math.Cos(th) * m
	y := math.Sin(ph) * math.Sin(th) * m
	z := math.Cos(ph) * m
	return Vec{x, y, z}
}

// Newp creates a new Vec vector from 2 points. The Vec represents the distance and angle from p1 to p1
func Newp(p1, p2 Vec) Vec {
	return p2.Minus(p1)
}

// Mul performs a scaler multiplation.
func (v Vec) Mul(x float64) Vec {
	m, theta, phi := v.AsSpherical()
	m *= x
	return Newd(m, theta, phi)
}

// Dot performs the dot product a.b
func (v Vec) Dot(b Vec) float64 {
	return v.X*b.X + v.Y*b.Y + v.Z*b.Z
}

// Cross performs the cross product u X v
func (v Vec) Cross(w Vec) Vec {
	x := v.Y*w.Z - v.Z*w.Y
	y := v.Z*w.X - v.X*w.Z
	z := v.X*w.Y - v.Y*w.X
	return Newc(x, y, z)
}
