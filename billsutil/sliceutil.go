// Package billsutil contains a set of miscellaneaous utilities Bill finds useful
// In particular it contains utilties for operating on slices using common functional idioms
package billsutil

import (
	"sort"
)

/* These are "mapping" functions common to collections classes in other languages
   I need t think about how to make these more elegant in the absence of generics in Go
   See: https://gobyexample.com/collection-functions for good examples
*/

// MapInt returns a new slice containing the results of applying the function f to each int in the original slice.
func MapInt(vs []int, f func(int) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// MapChar returns a new slice containing the results of applying the function f to each string in the original slice.
func MapChar(vs string, f func(rune) rune) string {
	str := ""
	for _, char := range vs {
		str += string(f(char))
	}
	return str
}

// ReduceInt returns the result of a pairwise application of the function through the slice
// The input slice must be at least length 2
func ReduceInt(vs []int, f func(int, int) int) int {
	if len(vs) < 1 {
		panic("Slice must be at least 1 element long")
	}
	if len(vs) == 1 {
		return vs[0] // This may be wrong! It's up to the caller to decide!
	}
	accum := vs[0]
	for i := 1; i < len(vs); i++ {
		accum = f(accum, vs[i])
	}
	return accum
}

// AllInt returns true if all of the ints in the slice satisfy the predicate f.
// The slice must be non empty
func AllInt(vs []int, f func(int) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// FindInt returns a slice of all positions of val in the slice vs. If there are no occurrences, an empty slice is returned.
func FindInt(vs []int, val int) []int {
	var result = []int{}
	for i, n := range vs {
		if n == val {
			result = append(result, i)
		}
	}
	return result
}

// Reverse reverses any sortable collection and is very efficient.
// Note that it reverse in place, so it modifies the slice passed in
// This is the implementation of the stringutil.Reverse example at https://github.com/golang/example
func Reverse(s sort.Interface) {
	len := s.Len()
	for front, back := 0, len-1; front < len/2; front, back = front+1, back-1 {
		s.Swap(front, back)
	}
}

// RuneSlice attaches the methods of sort.Interface to []rune, for sorting and reversing
type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ReverseStr is a convenience method for reversing strings
func ReverseStr(str string) string {
	r := RuneSlice([]rune(str))
	Reverse(r)
	return string(r)
}

// ReverseInt is a convenience method for reversing Int slices
// It makes a copy of the input to avoid mutating it.
func ReverseInt(is []int) []int {
	js := sort.IntSlice(append(is[:0:0], is...))
	Reverse(js)
	return js
}
