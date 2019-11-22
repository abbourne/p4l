package main

import (
	"fmt"
	"p4l/billsutil"
	"p4l/hw1"
)

func main() {

	fmt.Println(hw1.FactArray(10))
	fmt.Println(hw1.FibArray(10))
	fmt.Println(billsutil.MinArray([]int{42, 53, 9, 24, 38}))
	fmt.Println(hw1.GCDArray([]int{4, 8, 24}))

	/*
		println("Using CombDumb:")
		println("C(4,2) = ", p4l1.CombDumb(4, 2))
		println("C(8,4) = ", p4l1.CombDumb(8, 4))
		println("C(8,1) = ", p4l1.CombDumb(8, 1))
	*/

	/*
		println("Using p41l.Comb")
		println("C(4,2) = ", p41l.Comb(4, 2))
		println("C(8,4) = ", p41l.Comb(8, 4))
		println("C(8,1) = ", p41l.Comb(8, 1))
		println("C(1000,998) = ", p41l.Comb(1000, 998))
	*/

	/*
		println("Using PermDumb:")
		println("P(4,2) = ", p4l1.PermDumb(4, 2))
		println("P(8,4) = ", p4l1.PermDumb(8, 4))
		println("P(8,1) = ", p4l1.PermDumb(8, 1))
		// println("P(1000, 2) = ", p4l1.PermDumb(1000, 2)) // panic: runtime error: integer divide by zero
		//println("P(1000, 998) = ", p4l1.PermDumb(1000, 998)) // panic: runtime error: integer divide by zero
	*/

	/*
		println("Using Perm:")
		println("P(4,2) = ", p4l1.Perm(4, 2))
		println("P(8,4) = ", p4l1.Perm(8, 4))
		println("P(8,1) = ", p4l1.Perm(8, 1))
		println("P(1000, 2) = ", p4l1.Perm(1000, 2))
		// println("P(1000, 998) = ", p4l1.Perm(1000, 998)) // int64 overflows (answer is roughly =2.0119363𝐸+2567!)
	*/
}
