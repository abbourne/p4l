package hw1

import (
	"math"
	"math/bits"
)

// Use Package initiation to compute the prime sieve once
func init() {
	sieveOfPrimes = SieveOfEratosthenes(maxprime)
}

const maxprime = 33550336

var sieveOfPrimes []int

// IsPerfect1 JiPi's Version from discussion board (JiPi and Henk Westerink)
func IsPerfect1(n int) bool {
	sum := 1
	sqint := int(math.Sqrt(float64(n)) + 0.5)
	if sqint*sqint == n {
		sum += sqint
	}
	maxVal := math.Sqrt(float64(n))
	for k := 2; float64(k) < maxVal; k++ {
		if n%k == 0 {
			sum += k + n/k
			if sum > n {
				return false
			}
		}
	}
	return (sum == n)
}

/*
// Henk Westerin's Version from discussion board (JiPi and Henk Westerink)
Function IsPerfect(n int) returning bool {

	make slice divisors of lentgh 0

   append divisors with 1

   maxScan = n / 2; p = 2

   while p < maxScan {

	   if n % p == 0: {

		   append divisors with p

		   maxScan = n / p                                      // we already have all divisors beyond this number

		   append divisors with maxScan    // as n can be divided by p, n / p is also a divisor

	   }

   increment p by 1

   }

   if sum(divisors) == n {

	   return true

   } else {

	   return false

   }

}
*/

// IsPerfect based on John Cox's geometric progression algorithm
func IsPerfect(n int) bool {
	return n == SumOfFactors(n)
}

// SumOfFactors John D Cox's algorithm to compute the sum o factors
// computes sum by using geometric progression.  The final sum will also
// include the number N as a factor so that needs to be subtracted before returning
func SumOfFactors(origN int) int {
	// the overall formula is the product of all prime factors p occurring k times
	// (p^(k+1) - 1) / (p - 1)

	// ok, first initialize result with factors of two
	numZeros := bits.TrailingZeros(uint(origN))
	n := origN >> numZeros
	result := ((2 << numZeros) - 1)

	// ok, now account for odd numbers.
	sqrtN := int(math.Sqrt(float64(n)))
	for theOddNumber := 3; theOddNumber <= sqrtN; theOddNumber += 2 {
		if n%theOddNumber == 0 {
			sum := 1
			term := 1
			for (n % theOddNumber) == 0 {
				n /= theOddNumber
				term *= theOddNumber
				sum += term
			}
			result *= sum
			sqrtN = int(math.Sqrt(float64(n)))
		}
	}

	// ok, now don't forget the last prime factor if there is one
	if n > 1 {
		result *= 1 + n
	}
	return result - origN
}

// IsPerfect2 takes an integer n, and returns true if it is a perfect number.
// It also returns []int slice with all the factors of n
// Is Perfect uses a modified trial division algorithm to factorize the input,
// which is the most inefficient way of factoring. Some optimizations:
//  - We only test to sqrt(n), and then when we find a factor we also add n/factor as a second factor
//  - If a number does not divide n, we also know that all multiples of that number won't divide it either.
//    so we keep a slice of "notDivisibleby" values, and check against them before we do the trial division
func IsPerfect2(n int) (bool, []int) {
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
		if res := IsPerfect(i); res {
			result = append(result, i)
		}
	}
	return result

}
