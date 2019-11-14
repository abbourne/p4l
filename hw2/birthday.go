package hw2

import "math/rand"

// BirthdayParadox takes a number of people numPeople, a number of trials numTrials, and returns the %
// of time that at least 2 people have the same birthday
// See https://en.wikipedia.org/wiki/Birthday_problem
func BirthdayParadox(numPeople, numTrials int) float64 {
	sameBdayCnt := 0
	for trial := 0; trial < numTrials; trial++ {
		var bdays = []int{}
		for pers := 0; pers < numPeople; pers++ {
			bdays = append(bdays, rand.Intn(365)+1)
		}
		if HasRepeat(bdays) {
			sameBdayCnt++
		}
	}
	return float64(sameBdayCnt) / float64(numTrials)
}
