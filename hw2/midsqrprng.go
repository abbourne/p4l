package hw2

import (
	"strconv"
	"strings"
)

// CountNumDigits Take an integer and counts the number of digits. It ignores any minus sign
func CountNumDigits(i int) int {
	return len(strings.Trim(strconv.Itoa(i), "-"))
}
