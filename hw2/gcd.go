package hw2

import "p4l/billsutil"

// TrivialGCD is copied from Phillip's Go code files
func TrivialGCD(a, b int) int {
	d := 1
	m := billsutil.Min(a, b)
	for p := 1; p <= m; p++ {
		// remainder operation Remainder(n, k) is n%k (e.g., 14%3 = 1)
		if a%p == 0 && b%p == 0 { // if statement is only true if both are true
			// if I'm here, I know that the answer to both questions is "YES"
			d = p
		}
	}
	return d
}

// EuclidGCD is copied from Phillip's Go code files
func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}
