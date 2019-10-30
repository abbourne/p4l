package hw2_test

import (
	"math"
	"p4l/hw2"
	"testing"
)

var testPatternCountData = []struct {
	pattern  string // input
	text     string // input
	resCount int    // result
}{
	{"GCG", "GCGCG", 2},
	{"CG", "ACGTACGTACGT", 3},
	{"AAA", "AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGA", 4},
	{"TTT", "AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAAC", 3},
	{"ACT", "GGACTTACTGACGTACG", 2},
	{"A", "", 0},
	{"", "TCAGA", 0},
}

var testMaxDictData = []struct {
	testData map[string]int
	result   int
}{
	{map[string]int{"ACT": 3, "GTGA": 6, "TA": 2}, 6},
	{map[string]int{"x1213y": 12}, 12},
	{map[string]int{"adkfdjk": -4, "adskf": -3, "fjdk": -7}, -3},
	{map[string]int{}, math.MinInt64},
}

func TestPatternCount(t *testing.T) {
	t.Log("Run PatternCount tests")
	for _, tc := range testPatternCountData {
		count := hw2.PatternCount(tc.pattern, tc.text)
		t.Logf("PatternCount(%v, %v): result: count %d", tc.pattern, tc.text, count)
		if count != tc.resCount {
			t.Errorf("Result not equal to expected! PatternCount(%v, %v): result: %d", tc.pattern, tc.text, count)
		}

	}

}

func TestMaxDict(t *testing.T) {
	t.Log("Run MaxDict tests")
	for _, tc := range testMaxDictData {
		result := hw2.MaxDict(tc.testData)
		t.Logf("MaxDixt(%v): result: %d", tc.testData, result)
		if result != tc.result {
			t.Errorf("Result not equal to expected! PatternCount(%v): result: %d", tc.testData, result)
		}

	}

}
