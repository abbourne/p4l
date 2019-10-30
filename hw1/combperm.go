package p4l1

// PermDumb is a naive implementation to calculate the Permutation statistic.
// It will fail when n is greatr than about 20, as fact(n) will overflow a 64 bit int
func PermDumb(n int, k int) int {
	//defer TimeIt(time.Now(), "permDumb")
	return Fact(n) / Fact(n-k)
}

// CombDumb is a naive implementation to calculate the Permutation statistic.
// It will fail when n is greater than about 20, as fact(n) will overflow a 64 bit int
func CombDumb(n int, k int) int {
	//defer TimeIt(time.Now(), "combDumb")
	return Perm(n, k) / Fact(k)
}

// Perm calculates the Permutation statistic
// P(n,k) - represents the number of ways to choose k items from a list of n items, when the
//           ordering matters
//        - It is equal to the product: n*(n-1)*(n-2)*...*(n-k+1)
//        - This can also be written as P(n,k) = n!/(n-k!)
//        - See https://www.calculatorsoup.com/calculators/discretemathematics/permutations.php
//          for an online perutations calculator to check results
// This implementation calculates n*(n-1)*(n-2)*...*(n-k+1) to minimise issues with integer overflow
func Perm(n int, k int) int {
	//defer TimeIt(time.Now(), "Perm")
	result := 1
	for i := n; i >= n-k+1; i-- {
		// log.Println(i, result, n, k)
		result = result * i
	}
	return result
}

// Comb calculates the combination statistic
// C(n,k) - represents the number of ways to choose k items from a list of n items, when the
//           ordering does not matter
//        - This can also be written as C(n,k) = n!/(n-k!)*k!
//        - Or P(n,k)/k!
//        - This can be expanded to (n/1*(n-1)/2*(n-2)/3*...*(n-k+1)/k) While it may look like
//          remainders can result, as long as you perform the calculations strict;y from left to
//          to right, it actually works, with no remainders at any time.
//        - See https://www.calculatorsoup.com/calculators/discretemathematics/combinations.php
//          for an online combinations calculator to check results
//        - for hints on how to efficiently implement C(n,k) see:
//        	- https://www.mathbootcamps.com/counting-with-combinations/
//          - https://en.wikipedia.org/wiki/Combination
func Comb(n int, k int) int {
	//defer TimeIt(time.Now(), "comb")
	result := 1
	// This is tricky! In the expansion above, we have to check that i < k, so that the
	// denominators only go from 1 to k. And we have to check that i <= n-k, so that the
	// numerators only go from n to (n-k+1)
	for i := 1; i <= k && (i <= n-k); i++ {
		result = (result * (n - i + 1)) / i
	}
	return result
}
