package hw2

import (
	"strconv"
	"strings"
)

// CountNumDigits Take an integer and counts the number of digits. It ignores any minus sign
// An eklegant 1 line solution, but very inefficient!
func CountNumDigits(i int) int {
	return len(strings.Trim(strconv.Itoa(i), "-"))
}

// ComputePeriodLength takes a slice as inout and return the length (period) of any repeating sequence
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
