package p4l1_test

import (
	"p4l1"
	"testing"
)

var perfectTests = []struct {
	n       int  // input
	resBool bool //  expected result
}{
	{10, false},
	{6, true},
	{24, false},
	{28, true},
	{496, true},
	{498, false},
	{8128, true},
	{137438691328, true},
	{137438691329, false},
	{137438691330, false},
}

func TestIsPerfect(t *testing.T) {
	t.Log("Run IsPerfect tests")
	for _, tc := range perfectTests {
		isActual, sumActual := p4l1.IsPerfect(tc.n)
		t.Logf("IsPerfect(%d): result: %t, factors: %v", tc.n, isActual, sumActual)
		if isActual != tc.resBool {
			t.Errorf("Result not equal to expected! IsPerfect(%d): result: %t, factors: %v", tc.n, isActual, sumActual)
		}

	}

	if !true {

	}
}
