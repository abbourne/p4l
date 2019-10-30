package main

import (
	"billsutil"
	"fmt"
	"p4l1"
	"strconv"
	"time"
)

func timeFindPerfect(n int) {
	defer billsutil.TimeIt(time.Now(), "FindPerfect with n= "+strconv.Itoa(n))
	fmt.Println("Calling FindPerfect with limit", n)
	fmt.Println(p4l1.FindPerfect(n))
}

func main() {
	/*
		fmt.Println(p4l1.IsPerfect(6))
		fmt.Println(p4l1.IsPerfect(10))
		fmt.Println(p4l1.IsPerfect(28))
		fmt.Println(p4l1.IsPerfect(496))
		fmt.Println(p4l1.IsPerfect(498))

		result := p4l1.SieveOfEratosthenes(498)

		fmt.Println(result)
		fmt.Println(result[19], result[38], result[190])



			limits := []int{8130, 33550340, 8589869060, 137438691400}
			for _, i := range limits {
				timeFindPerfect(i)
			}
	*/

	fmt.Println("PrimeFactors 8128", p4l1.PrimeFactors(8128))
	fmt.Println("PrimeFactors 33550336", p4l1.PrimeFactors(33550336))

	fmt.Println("PrimeFactors2 8128", p4l1.PrimeFactors2(8128))
	fmt.Println("PrimeFactors2 33550336", p4l1.PrimeFactors2(33550336))

	fmt.Println(p4l1.IsPerfect(8128))
	fmt.Println(p4l1.IsPerfect(33550336))

}
