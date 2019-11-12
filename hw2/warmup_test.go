package hw2_test

import (
	"math/rand"
	"os"
	"p4l/billsutil"
	"p4l/hw2"
	"testing"
	"time"
)

// Note that this is used for *all* test files!
const LOG_ON = false

// Note that this is used for *all* test files!
func TestMain(m *testing.M) {
	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}
func TestWeightedDie(t *testing.T) {
	if LOG_ON {
		t.Log("Run WeightedDie tests")
	}
	trials := 100000
	results := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0}
	for i := 0; i < trials; i++ {
		results[hw2.WeightedDie()]++
	}
	var pOfi [6]float64
	for key, val := range results {
		pOfi[key-1] = float64(val) / float64(trials)
	}
	if LOG_ON {
		t.Log("Probabilities: ", pOfi)
	}
	if pOfi[2] < .45 || pOfi[2] > .55 {
		t.Error("Error probability of 3 is out of range. Got: ", pOfi[2])
	}
}

type trial struct {
	low, high int
}

var randPairTestTable = []trial{
	{1000, 2000},
	{20000, 10000},
	{100000, 200000},
	{2000000, 1000000}}

var numTrials = 10

func inRange(b1, b2, val int) bool {
	low := billsutil.Min(b1, b2)
	high := billsutil.Max(b1, b2)
	return b1 >= low && b2 <= high
}
func TestRandPairInRange(t *testing.T) {
	if LOG_ON {
		t.Log("Run RandPairInRange tests")
	}
	for _, trial := range randPairTestTable {
		for i := 0; i < numTrials; i++ {
			res1, res2 := hw2.RandPairInRange(trial.low, trial.high)
			if LOG_ON {
				t.Logf("From: %7d To: %7d  Res1: %7d Res2: %7d \n", trial.low, trial.high, res1, res2)
			}
			if !inRange(trial.low, trial.high, res1) || !inRange(trial.low, trial.high, res2) {
				t.Errorf("Returned value out of range - From: %7d To: %7d  Res1: %7d Res2: %7d \n",
					trial.low, trial.high, res1, res2)
			}
		}
	}
}
