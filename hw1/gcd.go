package p4l1

import "billsutil"

// TrivialGcd find the GCD of two ints in a very inefficient way
func TrivialGcd(a, b int) int {
	var lesser = billsutil.Min(a, b)
	var gcd = -1
	for i := 2; i <= lesser; i++ {
		if (a%i == 0) && (b%i == 0) {
			gcd = i
		}
	}
	return gcd
}

// GCDArray finds the GCD of slice of ints.
func GCDArray(nums []int) int {
	smallest := billsutil.MinArray(nums)
	gcd := -1
	for i := 2; i <= smallest; i++ {
		if billsutil.All(nums, func(v int) bool { return (v%i == 0) }) {
			gcd = i
		}
	}
	return gcd
}
