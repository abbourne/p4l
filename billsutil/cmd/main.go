package main

import (
	"fmt"
	"p4l/billsutil"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minArray(js []int) int {
	return billsutil.ReduceInt(js, min)
}

func main() {
	js := []int{42, 53, 9, 99, 24, 38, 0, -10, 2}
	fmt.Printf("Find Minimum and Maximum of slice %v \n", js)
	fmt.Println("The maximum is: ", billsutil.ReduceInt(js, max))
	fmt.Println("The minimum is: ", billsutil.ReduceInt(js, min))

	fmt.Println("The minimum is: ", billsutil.ReduceInt(js, func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}))
}
