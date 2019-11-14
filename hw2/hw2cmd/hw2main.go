package main

import (
	"fmt"
	"math/rand"
	"os"
	"p4l/hw2"
	"text/tabwriter"
	"time"
)

func main() {
	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())

	// Exercise 1: Compare Trivial and Euclid GCD Algorithms
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

	fmt.Println("Time for GCD:\n-----------------\n ")
	writer := tabwriter.NewWriter(os.Stdout, 10, 10, 1, ' ', 0)
	fmt.Fprintln(writer, "\t\t Trivial\tDuration\tDuration\t")
	fmt.Fprintln(writer, "   Low\t   High\tSlower by\tTrivial\t Euclid\t")
	writer.Flush()
	for _, t := range GCDTrialTable {
		t.durT = hw2.RunGCD(t.low, t.high, numTrials, hw2.TrivialGCD)
		t.durE = hw2.RunGCD(t.low, t.high, numTrials, hw2.EuclidGCD)
		diff := t.durT / t.durE
		fmt.Fprintf(writer, "%7d\t%7d\t%6d\t%v\t%v\t\n", t.low, t.high, diff, t.durT, t.durE)
	}
	writer.Flush()

	// Exercise 2: Relatively Prime Probability
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
	fmt.Println("\nRelatively Prime Probabilities:\n--------------------------\n ")
	writer = tabwriter.NewWriter(os.Stdout, 10, 10, 1, ' ', 0)
	fmt.Fprintln(writer, "    Low\t    High\tProbability\t")
	writer.Flush()
	for _, t := range RelativePrimeTrialTable {
		t.prob = hw2.RelativelyPrimeProbability(t.low, t.high, numTrials)
		fmt.Fprintf(writer, "%10d\t%10d\t%4f\t\n", t.low, t.high, t.prob)
	}
	writer.Flush()

	// Exercise 2: Birthday Paradox
	numTrials = 1000
	BdayTrialData := []int{5, 10, 15, 20, 23, 25, 30, 35, 40, 45, 50, 55, 60}
	fmt.Println("\nBirthday Paradox:/n------------------\n ")
	writer = tabwriter.NewWriter(os.Stdout, 10, 10, 1, ' ', 0)
	fmt.Fprintln(writer, " # of\t\t")
	fmt.Fprintln(writer, "People\tProbability\t")
	writer.Flush()
	for _, numPeople := range BdayTrialData {
		fmt.Fprintf(writer, "%3d\t%4f\t\n", numPeople, hw2.BirthdayParadox(numPeople, numTrials))
	}
	writer.Flush()

	fmt.Println("\nMiddle Square PRNG\n----------------------\n ")
	fmt.Println(hw2.GenerateMiddleSquareSequence(1600, 4))
	fmt.Println(hw2.GenerateMiddleSquareSequence(3792, 4))

	cnt := 0
	maxPeriod := 1
	for i := 1000; i < 10000; i++ {
		res := hw2.ComputePeriodLength(hw2.GenerateMiddleSquareSequence(i, 4))
		if res <= 10 {
			cnt++
		}
		if res > maxPeriod {
			maxPeriod = res
		}
	}
	fmt.Println("There are", cnt, "generators with a period of 10 or less")
	fmt.Println(maxPeriod, "is the largest period")

	fmt.Println("\nLinearCongruential PRNG\n--------------------\n ")

	res := hw2.GenerateLinearCongruentialSequence(1, 5, 1, 8192)
	fmt.Println("LinearCongruentialPRNG: seed=1, a=5, c=1, m=8192 has period:", hw2.ComputePeriodLength(res))

	cnt = 0
	m := 8191
	fmt.Printf("\nCounting LinearCongruentialPRNG (2,a,0,%d) with period of %d\n", m, m-1)
	for a := 1; a < m-1; a++ {
		res := hw2.ComputePeriodLength(hw2.GenerateLinearCongruentialSequence(2, a, 0, m))
		if res == m-1 {
			fmt.Printf("LinearCongruentialPRNG: seed=2, a= %d, c=0, m=%d has period: %d\n", a, m, res)
		}
	}
	fmt.Printf("found %d LinearCongruentialPRNGs (2,a,0,%d) with period of %d\n", cnt, m, m-1)

}
