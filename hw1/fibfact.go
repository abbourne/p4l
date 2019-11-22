package hw1

// Fact calculates the factorial of an int
// There are 2 big issues:
//  - the result grows quickly, and will soon overflow a 64 bit int at about n=20. One can solve
//    this by using bignums in the package math/big
//  - Go does not implement tail recursion, so each recursive call uses a stack frame. If the
//    recursion is very deep, we can run out of stack space. The solution is to use a for loop
//    to replace the recursion
func Fact(n int) int {
	if n < 2 {
		return 1
	}
	return n * Fact(n-1)
}

// FactArray makes an array of a sequence of n factorials
func FactArray(n int) []int {
	facts := make([]int, n)
	facts[0] = 1
	for i := 1; i < n; i++ {
		facts[i] = (i + 1) * facts[i-1]
	}
	return facts
}

// FibArray makes an array of a sequence of n Fibonacci numbers
func FibArray(n int) []int {
	fibs := make([]int, n)
	fibs[0] = 1
	fibs[1] = 1
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs
}
