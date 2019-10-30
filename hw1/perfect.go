package p4l1

import (
	"math"
)

// Use Package initiation to compute the prime sieve once
func init() {
	sieveOfPrimes = SieveOfEratosthenes(maxprime)
}

const maxprime = 33550336

var sieveOfPrimes []int

// IsPerfect takes an integer n, and returns true if it is a perfect number.
// It also returns []int slice with all the factors of n
// Is Perfect uses a modified trial division algorithm to factorize the input,
// which is the most inefficient way of factoring. Some optimizations:
//  - We only test to sqrt(n), and then when we find a factor we also add n/factor as a second factor
//  - If a number does not divide n, we also know that all multiples of that number won't divide it either.
//    so we keep a slice of "notDivisibleby" values, and check against them before we do the trial division
func IsPerfect(n int) (bool, []int) {
	bound := int(math.Sqrt(float64(n))) + 1
	isNotDivisibleBy := make([]bool, bound+1)
	factors := []int{1}
	sum := 1
	for i := 2; i < bound; i++ {
		if !isNotDivisibleBy[i] {
			if n%i == 0 {
				factors = append(factors, i, n/i)
				sum += i + n/i
			} else {
				for j := i; j < bound; j = j + i {
					isNotDivisibleBy[j] = true
				}
			}
		}
		if sum > n {
			return false, factors
		}
	}
	return sum == n, factors
}

//SieveOfEratosthenes takes an integer n and returns a slice of n+1 ints primeArray where primeArray[p] is 0 if p
// is prime and holds the lowest primer divisor of p if it is not
//It implements the Sieve of Eratosthenes approach.
func SieveOfEratosthenes(n int) []int {
	primeArray := make([]int, n+1)
	bound := int(math.Sqrt(float64(n)))

	// now, range over primeArray, and cross off multiples of first prime we see and iterate this process.
	for p := 2; p <= bound; p++ {
		if primeArray[p] == 0 {
			primeArray = CrossOffMultiples(primeArray, p)
		}
	}

	return primeArray
}

//CrossOffMultiples takes a boolean slice and an integer p.  It crosses off
// multiples of p, meaning turning these multiples to false in the slice.
// It returns the slice obtained as a result.
func CrossOffMultiples(primeArray []int, p int) []int {
	n := len(primeArray) - 1
	for k := 2 * p; k <= n; k += p {
		// all these multiples should be made composite
		primeArray[k] = p
	}
	return primeArray
}

// PrimeFactors returns all the prime factors of n
func PrimeFactors(n int) []int {
	//limit := int(math.Sqrt(float64(n)))
	tn := n
	factors := []int{1}
	for tn > 1 {
		if sieveOfPrimes[tn] == 0 {
			// tn is the smallest remaining prime factor
			factors = append(factors, tn)
			tn = 1 // will exit loop
		} else {
			// sieve[tn] is the smallest remaining prime factor
			factors = append(factors, sieveOfPrimes[tn])
			tn = tn / sieveOfPrimes[tn]
		}
	}
	return factors

}

// PrimeFactors2 uses simple trial division
func PrimeFactors2(n int) []int {
	factors := []int{1}
	f := 2 // First possible factor
	for n > 1 {
		if n%f == 0 {
			factors = append(factors, f)
			n /= f
		} else {
			f++
		}
	}
	return factors
}

// FindPerfect finds all perfect numbers up to some bound n. It returns a slice of the perfect numbers found
// Note that ISPerfect creates an array of length sqrt(n), so space as well as time will be a problem when n grows
// FindPerfect assumes perfect numbers are always even.
func FindPerfect(n int) []int {
	var result []int
	for i := 2; i <= n; i = i + 2 {
		if res, _ := IsPerfect(i); res {
			result = append(result, i)
		}
	}
	return result

}
