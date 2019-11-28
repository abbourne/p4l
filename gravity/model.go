package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"p4l/gifhelper"
	"p4l/vec"
	"strconv"
)

// Universe holds our bodies and the universe's parameters and
type Universe struct {
	bodies []Body
	width  float64
	g      float64 // Universal gravitational constant = 6.674E-11 Nm^2/kg^2
}

// Body represents any mass in our universe
type Body struct {
	name   string
	color  color.RGBA
	radius float64
	mass   float64
	pos    vec.Vec //Really a Point
	vel    vec.Vec
	accel  vec.Vec
}

// G is the universal gravitation constant, obtained from the initial Universe
var G float64

// SimulateGravity is our simulation engine. It takes an initial state, the size of the time interval for
// each "tick" and the number of simulation steps to run
func SimulateGravity(initialUniverse Universe, numGens int, t float64) []Universe {
	G = initialUniverse.g
	timePoints := make([]Universe, numGens+1)
	timePoints[0] = initialUniverse
	for i := 1; i < numGens; i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], t)
	}
	return timePoints
}

// UpdateUniverse takes a current universe, and returns a new universe calculated after time t has elapsed
func UpdateUniverse(curUni Universe, t float64) Universe {
	newUni := CopyUniverse(curUni)
	for i, body := range newUni.bodies {
		accel := ComputeNetAccel(curUni, body)
		// v2 = v1 + accl * t
		vel := body.vel.Plus(accel.Mul(t))
		// p2 = p1 + v*t + 1/2a*t^2
		pos := body.pos.Plus(vel.Mul(t)).Plus(accel.Mul(.5 * math.Pow(t, 2)))
		newUni.bodies[i].accel = accel
		newUni.bodies[i].vel = vel
		newUni.bodies[i].pos = pos
	}

	// need to update the body and put it into []body
	return newUni
}

// CopyUniverse creates a deepcopy of a universe. (values, not references)
func CopyUniverse(curUni Universe) (newUni Universe) {
	newUni.width = curUni.width
	newUni.g = curUni.g
	newUni.bodies = make([]Body, len(curUni.bodies))
	_ = copy(newUni.bodies, curUni.bodies)
	return
}

// ComputeNetAccel calculates the new acceleration on a body based on the gravitation effects of
// all other bodies in the universe
func ComputeNetAccel(curUni Universe, b Body) vec.Vec {
	netAccel := vec.Newc(0, 0, 0)
	for _, b2 := range curUni.bodies {
		if b2 != b {
			accel := ComputeAccl(b, b2)
			netAccel = netAccel.Plus(accel)
		}
	}
	return netAccel
}

// ComputeAccl compute the acceleration vector between two bodies
func ComputeAccl(b, b2 Body) vec.Vec {
	// Create a new vector which represents the distance and direction from b to b2
	// and get its spherical coordinates (distance and direction)
	dist, theta, phi := vec.Newp(b.pos, b2.pos).AsSpherical()
	accl := (G * b2.mass) / math.Pow(dist, 2)
	// Calculate the acceleration on b caused by b2
	// The acceleration is in the same direction as b2
	return vec.Newd(accl, theta, phi)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// WriteUniverse1 writes out the bodies in a Universe on a single line
func WriteUniverse1(resFile *os.File, i int, iUni Universe) {
	for _, body := range iUni.bodies {
		fmt.Fprintf(resFile, "%6d,%10s,%v,%v,%v\n", i, body.name, body.pos, body.vel, body.accel)
	}
}

func main() {

	// Jupiter System Model from Phillip
	jupiter := Body{
		name:   "Jupiter",
		color:  color.RGBA{223, 227, 202, 255},
		radius: 71000000,
		mass:   1.898e27,
		pos:    vec.Newc(2000000000, 2000000000, 0),
		vel:    vec.Newc(0, 0, 0),
		accel:  vec.Newc(0, 0, 0),
	}

	io := Body{
		name:   "Io",
		color:  color.RGBA{249, 249, 165, 255},
		radius: 1821000,
		mass:   8.9319e22,
		pos:    vec.Newc(2000000000-421600000, 2000000000, 0),
		vel:    vec.Newc(0, -17320, 0),
		accel:  vec.Newc(0, 0, 0),
	}

	europa := Body{
		name:   "Europa",
		color:  color.RGBA{132, 83, 52, 255},
		radius: 1569000,
		mass:   4.7998e22,
		pos:    vec.Newc(2000000000, 2000000000+670900000, 0),
		vel:    vec.Newc(-13740, 0, 0),
		accel:  vec.Newc(0, 0, 0),
	}

	ganymede := Body{
		name:   "Ganymede",
		color:  color.RGBA{76, 0, 153, 255},
		radius: 2631000,
		mass:   1.4819e23,
		pos:    vec.Newc(2000000000+1070400000, 2000000000, 0),
		vel:    vec.Newc(0, 10870, 0),
		accel:  vec.Newc(0, 0, 0),
	}

	callisto := Body{
		name:   "Callisto",
		color:  color.RGBA{0, 153, 76, 255},
		radius: 2410000,
		mass:   1.0759e23,
		pos:    vec.Newc(2000000000, 2000000000-1882700000, 0),
		vel:    vec.Newc(8200, 0, 0),
		accel:  vec.Newc(0, 0, 0),
	}

	jupiterSystem := Universe{
		width:  4000000000,
		bodies: []Body{jupiter, io, europa, ganymede, callisto},
	}
	/*  Earth/Moon Model
	Gravitational constant G = 6.67408e-11
	Name Colour Radius   Mass      PosX  PosY VelX VelY    AccX AccY
	Earth green 6.371e6  5.97237e24 -4.6e6 0.0 0.0 -12.14  0.0 0.0
	Moon grey   1.7374e6 7.342e22   3.80e8 0.0 0.0 1.022e3 0.0 0.0
	*/
	/*
		earth := Body{
			name:   "earth",
			color:  color.RGBA{0, 255, 0, 255}, // green
			radius: 6.371e6,
			mass:   5.97237e24,
			pos:    vec.Newc(-4.6e6, 0.0, 0),
			vel:    vec.Newc(0, -12.14, 0),
			accel:  vec.Newc(0, 0, 0),
		}

		moon := Body{
			name:   "moon",
			color:  color.RGBA{64, 64, 64, 255}, // grey,
			radius: 1.7374e6,
			mass:   7.342e22,
			pos:    vec.Newc(3.80e8, 0.0, 0),
			vel:    vec.Newc(0, 1.022e3, 0),
			accel:  vec.Newc(0, 0, 0),
		}

		earthSystem := Universe{
			bodies: []Body{earth, moon},
			width:  100000000,
			g:      6.67408e-11,
		}
	*/
	fmt.Println("Starting gravity simulator...")

	// numGens is in os.Args[1], time is in os.Args[2],
	// width of the drawing is in os.Args[3], os.Args[4] is to draw every nth image

	numGens, err1 := strconv.Atoi(os.Args[1])
	if err1 != nil {
		panic(err1)
	}
	time, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2 != nil {
		panic(err2)
	}
	canvasWidth, err3 := strconv.Atoi(os.Args[3])
	if err3 != nil {
		panic(err3)
	}
	frequency, err4 := strconv.Atoi(os.Args[4])
	if err4 != nil {
		panic(err4)
	}

	timePoints := SimulateGravity(jupiterSystem, numGens, time)
	fmt.Println("Simulation Run Successfully...")

	images := AnimateSystem(timePoints, canvasWidth, frequency)
	fmt.Println("Simulation Images Generated...")

	filename := "Universe_Animation"
	gifhelper.ImagesToGIF(images, filename)
	fmt.Println("Animation Generated...")

	resFileName := filename + ".txt"
	resFile, err := os.Create(resFileName)
	check(err)
	// Range over each universe
	fmt.Fprintln(resFile, "----START---")
	for i, iUni := range timePoints {
		if i%frequency == 0 {
			WriteUniverse1(resFile, i, iUni)
		}
	}
	fmt.Fprintln(resFile, "----END---")
	err = resFile.Close()
	check(err)

	fmt.Println("Data File Written... Done!")
}
