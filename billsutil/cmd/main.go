package main

import (
	"billsutil"
	"fmt"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func minArray(js []int) int {
	return billsutil.Reduce(js, min)
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(minArray([]int{42, 53, 9, 24, 38}))

}
