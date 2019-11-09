// Rules for craps and the odds https://www.playsmart.ca/table-games/craps/how-to-play
// and https://en.wikipedia.org/wiki/Craps
// The actual odds are -1.41% on the pass line and -1.36% on the don't pass line
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Play Craps")
	// Seed the PRNG
	rand.Seed(time.Now().UnixNano())
	numTrials := 1000000
	fmt.Println(ComputeHouseEdge(numTrials))
}

// math/rand common functions:
//  1. rand.Int: pseudorandom int
//  2. rand.Float64: pseudorandom decimal between [0, 1]
//  3. rand.Intn(n): pseudorandom int between 0 and n-1, inclusively

// RollDie simply returns the value of one die from 1-6 with equal probability
func RollDie() int {
	return rand.Intn(6) + 1
}

// IsWinner In craps and other casino games you can win, lose, or draw.
// A draw/tie is called a "push" and no monet chances hands
type IsWinner int

// In craps and other casino games you can win, lose, or draw.
// A draw/tie is called a "push" and no money chances hands
const (
	LOSE IsWinner = 0 + iota
	WIN
	PUSH
)

// SumTwoDice returns the sum of rolling two dice
func SumTwoDice() int {
	return RollDie() + RollDie()
}

// PlayCrapsOnce plays a full game of craps and returns results of 2 bets
// A "Pass" bet which is the most basic bet, and a "Don't Pass", which is almost the
// opposite of the Pass bet
func PlayCrapsOnce() (IsWinner, IsWinner) {
	comeOutRoll := SumTwoDice()
	switch comeOutRoll {
	case 7, 11:
		return WIN, LOSE
	case 2, 3:
		return LOSE, WIN
	case 12:
		return LOSE, PUSH
	}
	for true {
		pointRoll := SumTwoDice()
		if pointRoll == comeOutRoll {
			// Winner
			return WIN, LOSE
		} else if pointRoll == 7 {
			// Loser
			return LOSE, WIN
		}
	}
	panic("Unreachable code!")
	return PUSH, PUSH
}

// ComputeHouseEdge takes a number of trials, and retirns an estimate of the
// house edge of craps for both Pass and Don't Pass bets played over numTrials
func ComputeHouseEdge(numTrials int) (float64, float64) {
	passCnt := 0
	dontPassCnt := 0
	for i := 0; i < numTrials; i++ {
		passRes, dontPassRes := PlayCrapsOnce()
		if passRes == WIN {
			passCnt++
		}
		if passRes == LOSE {
			passCnt--
		}
		if dontPassRes == WIN {
			dontPassCnt++
		}
		if dontPassRes == LOSE {
			dontPassCnt--
		}
	}

	return float64(passCnt) / float64(numTrials), float64(dontPassCnt) / float64(numTrials)

}
