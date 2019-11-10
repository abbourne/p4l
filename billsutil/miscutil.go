package billsutil

import (
	"log"
	"math"
	"time"
)

// TimeIt is used as a defered function at the start of a function to time it
//   Usage:
//		func myFunction(n int) []bool {
//			defer TimeIt(time.Now(), "myFunction")
//			// rest of code goes here
//		}
func TimeIt(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
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

// Min returnd the minimum of two ints
func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// MinArray takes a []int and return the minimum element of the slice
func MinArray(js []int) int {
	return Reduce(js, Min)
}
