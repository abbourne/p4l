package hw2

import (
	"math/rand"
	"p4l/billsutil"
	"time"
)

// math/rand common functions:
//  1. rand.Int: pseudorandom int
//  2. rand.Float64: pseudorandom decimal between [0, 1]
//  3. rand.Intn(n): pseudorandom int between 0 and n-1, inclusively

// WeightedDie biases 3 to have a probability of .5
func WeightedDie() int {
	switch rand.Intn(10) {
	case 0, 1, 2, 3, 4:
		return 3
	case 5:
		return 1
	case 6:
		return 2
	case 7:
		return 4
	case 8:
		return 5
	case 9:
		return 6
	}
	return 0 // unreachable
}

// Return a random  int the range specified, inclusive of the requested limits.
// The limits may be specified in any order
func randIntInRange(from, to int) int {
	return rand.Intn(billsutil.Abs(to-from)+1) + billsutil.Min(to, from)
}

// RandPairInRange Returns a pair of unique random ints in the range specified inclusive of the requested limits.
// The limits may be specified in any order
func RandPairInRange(from, to int) (int, int) {
	i := randIntInRange(from, to)
	j := randIntInRange(from, to)
	for i == j {
		j = randIntInRange(from, to)
	}
	return i, j
}

// RunGCD runs a randomized test of a GCD function, and calculates the total time
// that the function takes, averaged over the total number of trials
func RunGCD(from, to, numTrials int, gcd func(a, b int) int) time.Duration {
	dur := time.Duration(0)
	for i := 1; i < numTrials; i++ {
		a, b := RandPairInRange(from, to)
		start := time.Now()
		_ = gcd(a, b)
		dur += (time.Now()).Sub(start)
		// log.Printf("gcd of %d, %d is %d", a, b, gcd)
	}
	return time.Duration(int(dur) / numTrials)
}
