// Package hw2 contains the answers to Programming for Lovers Homework 2 exercises.
// These exercises can be found on Stepik at https://stepik.org/course/59611/syllabus
// Definitions:
//  a DNA string or pattern is made up of the characters A, C, G, T representing each nucleotide
//  a k-mer is simply  DNA fragment of length k (i.e a substring)
//  an (L, t)-clump: given integers L and t, a string Pattern forms an (L, t)-clump inside a (larger)
//  string Genome if there is an interval of Genome of length L in which Pattern appears at least t times.
package hw2

import (
	"math"
	"strings"
)

// PatternMatching returns a slice of integers indicating the starting position of each occurrence of pattern in text
// If there are no occurrences, and empty slice is returned.
// Overlapping patterns are counted multiple times
func PatternMatching(pattern, text string) []int {
	result := []int{}
	k := len(pattern)
	n := len(text)
	if n == 0 || k == 0 || k > n {
		return result
	}
	for i := 0; i <= n-k; i++ {
		if text[i:i+k] == pattern {
			result = append(result, i)
		}
	}
	return result
}

// PatternCount takes a string text, and counts all occurances of pattern in text, including overlaps
// Note that this version only works with ASCII characters. It may mess up when Unicode characters are included
func PatternCount(pattern, text string) (count int) {
	k := len(pattern)
	n := len(text)
	if n == 0 || k == 0 || k > n {
		return
	}
	for i := 0; i <= n-k; i++ {
		if text[i:i+k] == pattern {
			count++
		}
	}
	return
}

// MaxDict finds the maximum of the int values in a map. It really does not care what the key is
func MaxDict(dict map[string]int) int {
	count := math.MinInt64
	for _, v := range dict {
		if v > count {
			count = v
		}
	}
	return count
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// FrequencyMap takes a text and returns all substrings (k-mers) of length k
func FrequencyMap(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)
	for i := 0; i <= n-k; i++ {
		pat := text[i : i+k]
		freqMap[pat]++
	}
	return freqMap
}

// FrequentWords takes a text, returns the substrings, (words, k-mers) of length k that have the most occurrences.
// More than one word may have the same maximum occurrences, so we return a slice.
func FrequentWords(text string, k int) (freqWords []string) {
	freqMap := FrequencyMap(text, k)
	maxCount := 0
	for _, count := range freqMap {
		if count > maxCount {
			maxCount = count
		}
	}
	for pat, count := range freqMap {
		if count == maxCount {
			freqWords = append(freqWords, pat)
		}
	}
	return
}

// Reverse takes an ASCII text string and reverses it.
// Note that this simplistic version will *not* work for strings wihh Unicode characters!
func Reverse(text string) string {
	text2 := ""
	for i := len(text) - 1; i >= 0; i-- {
		text2 += string(text[i])
	}
	return text2
}

// Complement  takes a DNA fragment as a text string and returns its complement
// The input should really only consist of ATCG characters. Other characters are passed unchanged
func Complement(pattern string) string {
	complements := map[string]string{"A": "T", "G": "C", "C": "G", "T": "A"}
	pattern = strings.ToUpper(pattern)
	converted := ""
	for i := 0; i < len(pattern); i++ {
		if comp, ok := complements[string(pattern[i])]; ok {
			converted += string(comp)
		} else {
			converted += string(pattern[i])
		}
	}

	return converted
}

// ReverseComplement takes a DNA fragment as a text string and returns its reverse complement.
// The input should only consist of ATCG characters
func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

// ClumpFinding takes a DNA string and finds all distinct k-mers forming (L, t)-clumps in the DNA.
func ClumpFinding(genome string, k, L, t int) []string {
	resultMap := map[string]int{}

	// Slide a windoe of length L through the genome
	var windowRes map[string]int
	for windowPos := 0; windowPos < len(genome)-L+1; windowPos++ {
		windowRes = FrequencyMap(genome[windowPos:windowPos+L], k)
		// Go through window res, find all results that occur t or more times, and add them to the result map
		for kmer, freq := range windowRes {
			if freq >= t {
				resultMap[kmer] += freq
			}
		}
	}
	// Extract the keys from the result map, and return them
	result := []string{}
	for key := range resultMap {
		result = append(result, key)
	}

	return result

}

// SkewArray takes a string representing a DNA genome as input and outputs a skew array of G-C
// We keep a counter. Every time we see a G we increment a the counter, every time we see a C we decrement it.
// For A and R we do nothing. We output a int array the same length as the genome, replacing the necleotide letters
// with the skew counter. We append the counter to the array first, *then* modify it.
func SkewArray(genome string) []int {
	var result = []int{}
	skew := 0
	for _, nucleotide := range genome {
		result = append(result, skew)
		switch nucleotide {
		case rune('G'):
			skew++
		case rune('C'):
			skew--
		}
	}
	result = append(result, skew)
	return result
}

// Reduce returns the result of a pairwise application of the function through the slice
// The input slice must be at least length 2
// NOTE: Copied from the package billsutil
func Reduce(vs []int, f func(int, int) int) int {
	accum := vs[0]
	for i := 1; i < len(vs); i++ {
		accum = f(accum, vs[i])
	}
	return accum
}

// Find returns a slice of all positions of  val in the slice vs. If there are no occances, an empty slice is returned.
func Find(vs []int, val int) []int {
	var result = []int{}
	for i, n := range vs {
		if n == val {
			result = append(result, i)
		}
	}
	return result
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// MinimumSkew generates the skew for a DNA genone string, then finds the positions of the minimum skew
func MinimumSkew(genome string) []int {
	var skewArray = SkewArray(genome)
	return Find(skewArray, Reduce(skewArray, min))
}
