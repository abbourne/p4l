package hw2

// IsRelativePrime returns true if the 2 ints are coprime, aka relatively prime
// See https://en.wikipedia.org/wiki/Coprime_integers and http://mathworld.wolfram.com/RelativelyPrime.html
func IsRelativePrime(i, j int) bool {
	if EuclidGCD(i, j) == 1 {
		return true
	}
	return false
}

// RelativelyPrimeProbability the relative probability that two ints selected form a range are co-prime
func RelativelyPrimeProbability(from, to, numTrials int) float64 {
	cnt := 0
	for i := 1; i < numTrials; i++ {
		if IsRelativePrime(RandPairInRange(from, to)) {
			cnt++
		}
	}
	return float64(cnt) / float64(numTrials)
}

