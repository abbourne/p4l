package hw2_test

import (
	"p4l/hw2"
	"testing"
)

type relPrimeTest struct {
	a, b int
	res  bool
}

var relPrimeTests = []relPrimeTest{
	{2, 4, false},
	{5, 6, true},
	{9, 12, false},
	{11, 13, true}}

func TestIsRelativePrime(t *testing.T) {
	if LOG_ON {
		t.Log("Run IsRelativePrime tests")
	}
	for _, trial := range relPrimeTests {
		res := hw2.IsRelativePrime(trial.a, trial.b)
		if res != trial.res {
			t.Errorf("IsRelativePrime Failed. Test Data: %v, %v, %v Result %v \n", trial.a, trial.b, trial.res, res)
		}

	}

}
