package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/golang/example/stringutil"
)

// NOTE: both reverseP and reverseB are quite inefficient since strings are immutable so
// text2 += string(text[i]) has to create a whole new copy of the string every time it is called!

// Phillips version
// reverse takes a string and reverses its symbols to produce a new string.
func reverseP(s string) string {
	s2 := ""
	n := len(s)

	for i := range s {
		s2 += string(s[n-1-i])
	}

	return s2
}

// Bills version
func reverseB(s string) string {
	s2 := ""
	for i := len(s) - 1; i >= 0; i-- {
		s2 += string(s[i])
	}
	return s2
}

// A not super efficient way of reversing a string that works for
// any UTF-8 string
func reverseRuneB(s string) string {
	s2 := ""
	for _, char := range s {
		s2 = string(char) + s2
	}
	return s2
}

// This version works for runes and is very efficient.
// W// We create the input string into an array of runes (which is mutable) and do the reverse in place
// We index from the front and the back, replace the 2 characters, then move 1 step toward the
// middle of the array from both ends
// This is the implementation of the stringutil.Reverse example at https://github.com/golang/example
// Its the Go idiomatic way to reverse a string
func reverse(s string) string {
	r := []rune(s) //convert from a string to an array of runes. Go will "unpack" the UTF8 into runes
	for front, back := 0, len(r)-1; front < len(r)/2; front, back = front+1, back-1 {
		r[front], r[back] = r[back], r[front]
	}
	return string(r) // And convert from a []rune array back to a UTF8 string.
}

// complementChar take a single character representing a nucleotide (ATGC) and
// returns it's complement. Any character other than ATGC is returned unchanged
func complementChar(char rune) rune {
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

	converted := ""
	for _, char := range pattern {
		converted += string(complementChar(char))
	}

	return converted
}

// Phillips version of complement
func complementP(dna string) string {
	dna2 := ""
	for i := range dna {
		switch dna[i] {
		case 'A':
			dna2 += "T"
		case 'T':
			dna2 += "A"
		case 'C':
			dna2 += "G"
		case 'G':
			dna2 += "C"
		}
	}
	return dna2
}

//ReverseComplement takes a DNA string and returns its reverse complement
//(corresponding to opposing strand).
func reverseComplement(dna string) string {
	return reverse(complement(dna))
}

// Go string formatting: https://gobyexample.com/string-formatting
func main() {
	s1 := "ATGC"
	fmt.Println("Input string:", s1, "is length: ", len(s1))
	fmt.Println("ReversedP:", reverseP(s1))
	fmt.Println("ReversedB:", reverseB(s1))

	s2 := "Hello, 世界" //copied from https://golang.org
	fmt.Println("Input string:", s2, "is length:", len(s2))
	fmt.Println("ReversedP:", reverseP(s2))
	fmt.Println("ReversedB:", reverseB(s2))

	s4 := string('⌘') + string('A') + string('𐅻') // encodes rune characters into UTF-8 strings
	fmt.Println(s4)
	s4runes := []rune(s4)     // decodes UTF-8 string into a slice of runes
	s4UTF8 := string(s4runes) // encodes a slice of runes back into a UTF-8 string
	fmt.Println(s4runes)
	fmt.Println(s4UTF8)

	s3 := "Hi 𐅻世⌘"
	fmt.Println("Input string:", s3, "is length:", len(s3))
	fmt.Println("Input string:", s3, "is runelength:", utf8.RuneCountInString(s3))
	for index, runeValue := range s3 {
		fmt.Printf("rune %#U has type %T starts at byte position %d\n", runeValue, runeValue, index)
		fmt.Printf("byte %#U has type %T starts at byte position %d\n", s3[index], s3[index], index)
	}

	maxChar := unicode.MaxRune
	fmt.Printf("%#U is the largest rune (code point) in Go\n", maxChar)

	if s1[0] == 'A' {
		fmt.Println("Complement of A is T")
	}
	/*
		var test1 uint8 = 65
		var test2 int32 = 8984
		test2 = test1  // This gives a type error, even though a uint8 will fit in an int32
	*/

	fmt.Println("Input string:", s2, "is length:", len(s2))
	fmt.Println("Reversed:", reverse(s2))
	fmt.Println("ReversedRuneB:", reverse(s2))
	fmt.Println("Reversed stringutil:", stringutil.Reverse(s2))

}
