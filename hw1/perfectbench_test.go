package hw1_test

import (
	"p4l/hw1"
	"testing"
)

var result bool
var factors []int

func benchmarkIsPerfect(n int, b *testing.B) {
	b.Logf("Run IsPerfect benchmark. n= %v", n)
	var res bool
	for i := 0; i < b.N; i++ {
		res = hw1.IsPerfect(n)
	}
	result = res
}

func benchmarkFindPerfect(n int, b *testing.B) {
	b.Logf("Run FindPerfect benchmark. n= %v", n)
	var res []int
	for i := 0; i < b.N; i++ {
		res = hw1.FindPerfect(n)
	}
	factors = res
}

func benchmarkPrimeFactors(n int, b *testing.B) {
	b.Logf("Run PrimeFactors benchmark. n= %v", n)
	var res []int
	for i := 0; i < b.N; i++ {
		res = hw1.PrimeFactors(n)
	}
	factors = res
}

func benchmarkPrimeFactors2(n int, b *testing.B) {
	b.Logf("Run PrimeFactors2 benchmark. n= %v", n)
	var res []int
	for i := 0; i < b.N; i++ {
		res = hw1.PrimeFactors2(n)
	}
	factors = res
}

//func BenchmarkIsPerfect1(b *testing.B)     { benchmarkIsPerfect(8128, b) }
func BenchmarkIsPerfect2(b *testing.B) { benchmarkIsPerfect(33550336, b) }

//func BenchmarkPrimeFactors1(b *testing.B)  { benchmarkPrimeFactors(8128, b) }
func BenchmarkPrimeFactors2(b *testing.B) { benchmarkPrimeFactors(33550336, b) }

//func BenchmarkPrimeFactors21(b *testing.B) { benchmarkPrimeFactors2(8128, b) }
func BenchmarkPrimeFactors22(b *testing.B) { benchmarkPrimeFactors2(33550336, b) }

/* func BenchmarkIsPerfect2(b *testing.B)   { benchmarkIsPerfect(33550336, b) }
func BenchmarkIsPerfect3(b *testing.B)   { benchmarkIsPerfect(137438691328, b) }
func BenchmarkIsPerfect4(b *testing.B)   { benchmarkIsPerfect(137438691329, b) }
func BenchmarkIsPerfect5(b *testing.B)   { benchmarkIsPerfect(137438691330, b) }
func BenchmarkFindPerfect1(b *testing.B) { benchmarkFindPerfect(500, b) }
func BenchmarkFindPerfect2(b *testing.B) { benchmarkFindPerfect(8130, b) }
*/

func BenchmarkSieveOfEratosthenes(b *testing.B) {
	n := 35000000
	var res []int
	b.Logf("Run SieveOfEratosthenes n= %v", n)
	for i := 0; i < b.N; i++ {
		res = hw1.SieveOfEratosthenes(n)
	}
	factors = res

}
