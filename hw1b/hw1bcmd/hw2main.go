package main

import (
	"fmt"
	"p4l/hw1b"
)

func main() {
	fmt.Println("Samples for Stepik String Functions Assignment ")

	fmt.Println(hw1b.ClumpFinding("ACGTACGT", 1, 5, 2))
	fmt.Println(hw1b.SkewArray("TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGAT"))

}
