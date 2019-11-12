package billsutil

import (
	"math"
	"time"
)

// TimeIt is used as a defered function at the start of a function to time it
//   Usage:
//		func myFunction(n int) []bool {
//			defer TimeIt(time.Now(), "myFunction")
//			// rest of code goes here
//		}
func TimeIt(start time.Time, name string) time.Duration {
	elapsed := time.Since(start)
	//log.Printf("%s took %s", name, elapsed)
	return elapsed
}

// TimedFunc takes any finction that does not return a value and times it
// The time taken is returned as a time.Duration
func TimedFunc(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// IsPrime is an inefficient test for whether an integer is prime or not
func IsPrime(p int) bool {
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k == 0 {
			return false
		}
	}
	return true
}

// Abs returns the absolute value of an int
func Abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

// Min returns the minimum of two ints
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// Max returns the maximum of two ints
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// MinArray takes a []int and return the minimum element of the slice
func MinArray(js []int) int {
	return ReduceInt(js, Min)
}
