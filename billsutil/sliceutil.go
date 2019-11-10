// Package billsutil contains a set of miscellaneaous utilities Bill finds useful
// In particular it contains utilties for operating on slices using common functional idioms
package billsutil

/* These are "mapping" functions common to collections classes in other languages
   I need t think about how to make these more elegant in the absence of generics in Go
   See: https://gobyexample.com/collection-functions for good examples
*/

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Reduce returns the result of a pairwise application of the function through the slice
// The input slice must be at least length 2
func Reduce(vs []int, f func(int, int) int) int {
	accum := vs[0]
	for i := 1; i < len(vs); i++ {
		accum = f(accum, vs[i])
	}
	return accum
}

// All returns true if all of the ints in the slice satisfy the predicate f.
// The slice must be non empty
func All(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Find returns a slice of all positions of val in the slice vs. If there are no occurrences, an empty slice is returned.
func Find(vs []int, val int) []int {
	var result = []int{}
	for i, n := range vs {
		if n == val {
			result = append(result, i)
		}
	}
	return result
}
