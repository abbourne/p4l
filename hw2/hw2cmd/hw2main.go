package main

import (
	"fmt"
	"math/rand"
	"p4l/hw2"
	"time"
)

func main() {
	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())

	numTrials := 100

	type GCDTrial struct {
		low, high  int
		durT, durE time.Duration
	}
	GCDTrialTable := []GCDTrial{
		{1000, 2000, 0, 0},
		{10000, 20000, 0, 0},
		{100000, 200000, 0, 0},
		{1000000, 2000000, 0, 0}}

	fmt.Println("Time for GCD: ")
	for _, t := range GCDTrialTable {
		t.durT = hw2.RunGCD(t.low, t.high, numTrials, hw2.TrivialGCD)
		t.durE = hw2.RunGCD(t.low, t.high, numTrials, hw2.EuclidGCD)
		diff := t.durT / t.durE
		fmt.Printf("Low: %7d, High: %7d: Trivial %6d times slower, TrivialGCD time: %v EuclidGCD time: %v \n",
			t.low, t.high, diff, t.durT, t.durE)
	}

	type RelativePrimeTrial struct {
		low, high int
		prob      float64
	}
	RelativePrimeTrialTable := []RelativePrimeTrial{
		{2, 1000, 0},
		{1000, 2000, 0},
		{2000, 3000, 0},
		{3000, 4000, 0},
		{4000, 5000, 0},
		{2, 10000, 0},
		{10000, 20000, 0},
		{20000, 30000, 0},
		{30000, 40000, 0},
		{40000, 50000, 0},
		{2, 100000, 0},
		{10000, 200000, 0},
		{2, 1000000, 0},
		{1000000, 2000000, 0},
		{2, 100000000, 0},
		{100000000, 1000000000, 0}}

	numTrials = 100000
	// Note that calculated probability is a little less than 61% (https://en.wikipedia.org/wiki/Coprime_integers)
	fmt.Println("\n Relatively Prime Probabilities: ")
	for _, t := range RelativePrimeTrialTable {
		t.prob = hw2.RelativelyPrimeProbability(t.low, t.high, numTrials)
		fmt.Printf("Low: %10d, High: %10d  Probability: %4f \n",
			t.low, t.high, t.prob)
	}
}
