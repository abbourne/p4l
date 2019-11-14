package hw2

import (
	"p4l/billsutil"
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

// Contains tells whether a contains x.
// Has Repeat is very slow. When creating a sequence, simply checking whether the
// element being added to the sequence already exists is enough to detect a cycle.
func Contains(is []int, j int) bool {
	for _, n := range is {
		if j == n {
			return true
		}
	}
	return false
}

// CountNumDigitss Take an integer and counts the number of digits. It ignores any minus sign
// An elegant 1 line solution, but very inefficient!
func CountNumDigitss(i int) int {
	return len(strings.Trim(strconv.Itoa(i), "-"))
}

// TakeMiddlens Takes the middle n digits of val.
// We do the manipulation with a string
func TakeMiddlens(val, n int) int {
	vals := strconv.Itoa(val)
	if n < 1 {
		panic("must take 1 or more digits")
	}
	//if n is odd, val must have an odd number of digits. If n is even, lenVal must be even
	if n%2 == 0 {
		if len(vals)%2 == 1 {
			vals = "0" + vals
		}
	} else {
		if len(vals)%2 == 0 {
			vals = "0" + vals
		}
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

// CountNumDigits Take an integer and counts the number of digits. It ignores any minus sign
func CountNumDigits(i int) int {
	i = billsutil.Abs(i)
	if i == 0 {
		return 1
	}
	cnt := 0
	for ; i > 0; cnt++ {
		i = i / 10
	}
	return cnt
}

// Pow10 returns an int of 10^i
func Pow10(i int) int {
	res := 1
	for ; i >= 1; i-- {
		res *= 10
	}
	return res
}

// TakeMiddlen Takes the middle n digits of val.
func TakeMiddlen(val, n int) int {
	if n < 1 {
		panic("must take 1 or more digits")
	}
	lenVal := CountNumDigits(val)

	//if n is odd, val must have an odd number of digits. If n is even, lenVal must be even
	if n%2 == 0 {
		if lenVal%2 != 0 {
			lenVal++
		}
	} else {
		if lenVal%2 == 0 {
			lenVal++
		}
	}
	trim := (lenVal - n) / 2
	val = val % Pow10(trim+n)
	val = val / Pow10(trim)
	return val

}

// MiddleSquare squares the input and then takes the numDigits digits from the middle of the result
// The input must contain an even number of digits
func MiddleSquare(seed, numDigits int) int {
	if numDigits < 2 || numDigits > 10 || numDigits%2 != 0 {
		panic("numdigits must be and event int between 2 and 10")
	}
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
	done := false
	for !done {
		val := prng()
		done = Contains(rs, val)
		rs = append(rs, val)
	}
	return rs
}
