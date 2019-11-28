package main

import "math"

func SimulateGravity(initialUniverse Universe, numGens int, time float64) []Universe {
	timePoints := make([]Universe, numGens+1)
	timePoints[0] = initialUniverse
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], time)
	}

	return timePoints
}

func UpdateUniverse(currentUniverse Universe, t float64) Universe {
	newUniverse := CopyUniverse(currentUniverse)
	for i := range newUniverse.bodies {
		newUniverse.bodies[i].acceleration = UpdateAcceleration(newUniverse.bodies, newUniverse.bodies[i])
		newUniverse.bodies[i].velocity = UpdateVelocity(newUniverse.bodies[i], t)
		newUniverse.bodies[i].position = UpdatePosition(newUniverse.bodies[i], t)
	}

	return newUniverse
}

func CopyUniverse(currentUniverse Universe) Universe {
	var newUniverse Universe
	newUniverse.width = currentUniverse.width
	newUniverse.bodies = make([]Body, len(currentUniverse.bodies))
	for i := range currentUniverse.bodies {
		var b Body
		currBody := currentUniverse.bodies[i]
		b.name = currBody.name
		b.red = currBody.red
		b.green = currBody.green
		b.blue = currBody.blue
		b.mass = currBody.mass
		b.radius = currBody.radius
		b.acceleration.x = currBody.acceleration.x
		b.acceleration.y = currBody.acceleration.y
		b.velocity.x = currBody.velocity.x
		b.velocity.y = currBody.velocity.y
		b.position.x = currBody.position.x
		b.position.y = currBody.position.y

		newUniverse.bodies[i] = b
	}
	return newUniverse
}

func UpdateAcceleration(bodies []Body, b Body) OrderedPair {
	var accel OrderedPair
	force := ComputeNetForce(bodies, b)
	accel.x = force.x / b.mass
	accel.y = force.y / b.mass
	return accel
}

func ComputeNetForce(bodies []Body, b Body) OrderedPair {
	var netForce OrderedPair
	for i := range bodies {
		if bodies[i] != b {
			force := ComputeForce(bodies[i], b)
			netForce.x += force.x
			netForce.y += force.y
		}
	}
	return netForce
}

func ComputeForce(b2, b Body) OrderedPair {
	var force OrderedPair
	d := Distance(b2.position, b.position)
	F := G * b2.mass * b.mass / (d * d)
	deltaX := b2.position.x - b.position.x
	deltaY := b2.position.y - b.position.y
	force.x = F * (deltaX / d)
	force.y = F * (deltaY / d)
	return force
}

func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func UpdateVelocity(b Body, time float64) OrderedPair {
	var v OrderedPair
	v.x = b.velocity.x + b.acceleration.x*time
	v.y = b.velocity.y + b.acceleration.y*time
	return v
}

func UpdatePosition(b Body, time float64) OrderedPair {
	var p OrderedPair
	p.x = b.position.x + b.velocity.x*time + .5*b.acceleration.x*time*time
	p.y = b.position.y + b.velocity.y*time + .5*b.acceleration.y*time*time
	return p
}
