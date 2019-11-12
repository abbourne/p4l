package main

import (
	"fmt"
	"p4l/billsutil"
	"unicode"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minArray(js []int) int {
	return billsutil.ReduceInt(js, min)
}

// reverse is a not super efficient way of reversing a string that works for
// any UTF-8 string
func reverse(s string) string {
	s2 := ""
	for _, char := range s {
		s2 = string(char) + s2
	}
	return s2
}

// ComplementChar take a single character representing a nucleotide (ATGC) and
// returns it's complement. Any character other than ATGC is returned unchanged
func ComplementChar(char rune) rune {
	switch unicode.ToUpper(char) {
	case 'A':
		return 'T'
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	default:
		return char
	}
}

// complement  takes a DNA fragment as a text string and returns its complement
// The input should really only consist of ATCG characters. Other characters are passed unchanged
func complement(pattern string) string {
	return billsutil.MapChar(pattern, ComplementChar)
}

// ReverseComplement takes a DNA fragment as a text string and returns its reverse complement.
// The input should only consist of ATCG characters
func reverseComplement(pattern string) string {
	return reverse(complement(pattern))
}

func main() {
	js := []int{42, 53, 9, 99, 24, 38, 0, -10, 2}
	fmt.Printf("Find Minimum and Maximum of slice %v \n", js)
	fmt.Println("The maximum is: ", billsutil.ReduceInt(js, max))
	fmt.Println("The minimum is: ", billsutil.ReduceInt(js, min))

	fmt.Println("The minimum is: ", billsutil.ReduceInt(js, func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}))

	patt := "AGTC"
	fmt.Println("The reverse complement of ", patt, " is ", reverseComplement(patt))

	fmt.Println("Playing with reverse ")
	fmt.Println("testInt reversed: ", billsutil.ReverseInt([]int{1, 2, 3, 4, 5}))
	fmt.Println("testStr reversed: ", billsutil.ReverseStr("世 ABCDE 𐅻 ⌘"))

}
