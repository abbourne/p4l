package main

import (
	"fmt"
	"p4l/billsutil"
	"p4l/hw1"
	"strconv"
	"time"
)

func timeFindPerfect(n int) {
	defer billsutil.TimeIt(time.Now(), "FindPerfect with n= "+strconv.Itoa(n))
	fmt.Println("Calling FindPerfect with limit", n)
	fmt.Println(hw1.FindPerfect(n))
}

func main() {
	fmt.Println(hw1.IsPerfect(6))
	fmt.Println(hw1.IsPerfect(10))
	fmt.Println(hw1.IsPerfect(28))
	fmt.Println(hw1.IsPerfect(496))
	fmt.Println(hw1.IsPerfect(498))
	fmt.Println(hw1.IsPerfect(8128))
	fmt.Println(hw1.IsPerfect(33550336))
	fmt.Println(hw1.IsPerfect(8589869056))

	fmt.Println(hw1.SumOfFactors(6))
	fmt.Println(hw1.SumOfFactors(33550336))

	limits := []int{8130, 33550340, 8589869060, 137438691400}
	for _, i := range limits {
		timeFindPerfect(i)
	}

	/*
		result := p4l1.SieveOfEratosthenes(498)

		fmt.Println(result)
		fmt.Println(result[19], result[38], result[190])

		fmt.Println("PrimeFactors 8128", hw1.PrimeFactors(8128))
		fmt.Println("PrimeFactors 33550336", hw1.PrimeFactors(33550336))

		fmt.Println("PrimeFactors2 8128", hw1.PrimeFactors2(8128))
		fmt.Println("PrimeFactors2 33550336", hw1.PrimeFactors2(33550336))

	*/

}
