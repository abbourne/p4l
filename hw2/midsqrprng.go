package hw2

import (
	"strconv"
	"strings"
)

// Sqrt of 2^64 is 4.3x10^9 (10 digits). 2^64 is 1.8x10^19 (20 digits)
// See https://en.wikipedia.org/wiki/Middle-square_method
// See https://en.wikipedia.org/wiki/Linear_congruential_generator

// ComputePeriodLength takes a slice as input and returns the length (period) of the first repeating sequence
// If no repeats are found, it returns 0
func ComputePeriodLength(is []int) int {
	for i := 0; i < len(is)-1; i++ {
		jc := is[i]
		for j := i + 1; j < len(is); j++ {
			if jc == is[j] {
				return j - i
			}
		}

	}
	return 0
}

// HasRepeat returns true if any of the elements in the slice are the same
func HasRepeat(is []int) bool {
	if ComputePeriodLength(is) == 0 {
		return false
	}
	return true
}

// CountNumDigits Take an integer and counts the number of digits. It ignores any minus sign
// An eklegant 1 line solution, but very inefficient!
func CountNumDigits(i int) int {
	return len(strings.Trim(strconv.Itoa(i), "-"))
}



// TakeMiddlen Takes the middle n digits of val.
func TakeMiddlen(val, n int) int {
	vals := strconv.Itoa(val)
	if n < 1 {
		panic("must take 1 or more digits")
	}
	//if n is odd, val must have an odd number of digits}
	if len(vals)%n != 0 {
		vals = "0" + vals
	}
	// pad with leading 0s if needed
	if pad := n - len(vals); pad > 0 {
		vals = strings.Repeat("0", pad) + vals
	}
	if len(vals) != n {
		trim := (len(vals) - n) / 2
		vals = string(([]rune(vals))[trim : n+trim])
	}
	res, _ := strconv.Atoi(vals)
	return res
}

// MiddleSquare squares the input and then takes the numDigits digits from the middle of the result
// The input must contain an even number of digits
func MiddleSquare(seed, numDigits int) int {
	if numDigits < 2 || numDigits > 10 || numDigits%2 != 0 {
		panic("numdigits must be and event int between 2 and 10")
	}
	/*
		if seed < 10 || CountNumDigits(seed)%2 != 0 {
			panic("seed must have an even number of digits and be positive")
		}
	*/
	return TakeMiddlen(seed*seed, numDigits)
}

// NewMiddleSquarePrng returns a new MiddleSquare random number generator function
func NewMiddleSquarePrng(seed, numDigits int) func() int {
	val := seed
	return func() int {
		val = MiddleSquare(val, numDigits)
		return val
	}
}

// GenerateMiddleSquareSequence returns an []int composed of generated random numbers
// It finishes as soon as a repeat is detected
func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	rs := []int{seed}
	prng := NewMiddleSquarePrng(seed, numDigits)
	for !HasRepeat(rs) {
		rs = append(rs, prng())
	}
	return rs
}

// LinearCongruential return the next value of a Linear Conruential Generator
// Where Xnext <- (a*X +c) mod m
func LinearCongruential(seed, a, c, m int) int {
	return (seed*a + c) % m
}

// NewLinearCongruentialPRNG returns a new Linear Conruential Generator random number generator function
// Where Xnext <- (a*X +c) mod m
func NewLinearCongruentialPRNG(seed, a, c, m int) func() int {
	val := seed
	return func() int {
		val = LinearCongruential(val, a, c, m)
		return val
	}
}

// GenerateLinearCongruentialSequence returns an []int composed of generated random numbers
// It finishes as soon as a repeat is detected
func GenerateLinearCongruentialSequence(seed, a, c, m int) []int {
	rs := []int{seed}
	prng := NewLinearCongruentialPRNG(seed, a, c, m)
	for !HasRepeat(rs) {
		rs = append(rs, prng())
	}
	return rs
}
