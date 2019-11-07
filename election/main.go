package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Let's simulate an election!")
	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())
	numTrials := 10000
	marginOfError := 0.10
	fileName := "debates.txt"

	// Read in the 2 files we are interested in
	electoralVotes := ReadElectoralVotes("electoralVotes.txt")
	polls := ReadPollingData(fileName)

	probability1, probability2, probabilityTie := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)

	fmt.Println("Estimated probability of ")
	fmt.Println("   Candidate1 win:", probability1)
	fmt.Println("   Candidate2 win:", probability2)
	fmt.Println("   Tie:           ", probabilityTie)

}

// SimulateMultipleElections runs our simulation
// Takes pooling data as a map of state names to %support as a float. (for
// candidate 1 - candidate 2 is just 100% - the %support for candidate 1), along with a map
// of state names to electoral college votes, a number of trials, and a marigin of error in the
// polls
// It returns the probabilities of candidate1, candidatew winning or a tie
func SimulateMultipleElections(
	polls map[string]float64,
	electoralVotes map[string]uint,
	numTrials int,
	marginOfError float64) (float64, float64, float64) {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	// simulate an election numtrials times and update the counts each time
	for i := 0; i < numTrials; i++ {
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)
		if votes1 > votes2 {
			winCount1++
		} else if votes1 < votes2 {
			winCount2++
		} else {
			tieCount++
		}
	}

	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	probabilityTie := float64(tieCount) / float64(numTrials)

	return probability1, probability2, probabilityTie
}

// SimulateOneElection runs a single election and returns the votes for each candidate
func SimulateOneElection(polls map[string]float64,
	electoralVotes map[string]uint,
	marginOfError float64) (uint, uint) {

	var collegeVotes1 uint = 0
	var collegeVotes2 uint = 0

	// range over all the states, and simulate the elction in each state
	for state := range polls {
		poll := polls[state] // The polling value in the state for candidate1
		numVotes := electoralVotes[state]
		// adjust the polling value for randomness based on margin of error
		adjustedPoll := AddNoise(poll, marginOfError)
		// Check on who won the state based on the adjustedPoll
		if adjustedPoll >= .5 {
			// candidate1 wins
			collegeVotes1 += numVotes
		} else {
			// candidate2 wins
			collegeVotes2 += numVotes
		}
	}
	return collegeVotes1, collegeVotes2
}

// AddNoise takes the polling value for candidate1, the polls margin of error, and returns an
// adjusted polling value after adding random noise
func AddNoise(poll float64, marginOfError float64) float64 {
	x := rand.NormFloat64() //random number from standard normal distribution
	x = x / 2               // x has a 95% chance of being between -1 and 1
	x = x * marginOfError   // x has a 95% chance of being between -marginOfError and marginOfError
	return x + poll
}
