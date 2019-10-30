package hw2

import "math"

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

// FrequentWords takes a text, returns the substrings, (words, k-mers) of length k that have the most occurances.
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
