package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(text string) string {
	text2 := ""
	for i := len(text) - 1; i >= 0; i-- {
		text2 += string(text[i])
	}
	return text2
}

// Go string formatting: https://gobyexample.com/string-formatting
func main() {
	s1 := "ABCDE"
	fmt.Println("Input string:", s1, "is length:", len(s1))
	fmt.Println("Input string:", s1, "is runelength:", utf8.RuneCountInString(s1))
	fmt.Println("Reversed:", reverse(s1))

	s1 = "Hello, ⌘世界" //copied from https://golang.org
	fmt.Println("Length of: H ", utf8.RuneLen('H'))
	fmt.Println("Length of: ⌘ ", utf8.RuneLen('⌘'))
	fmt.Println("Input string:", s1, "is length:", len(s1))
	fmt.Println("Input string:", s1, "is runelength:", utf8.RuneCountInString(s1))
	fmt.Println("Reversed:", reverse(s1))

	for index, runeValue := range s1 {
		fmt.Printf("%#U has type %T starts at byte position %d\n", runeValue, runeValue, index)
	}

	strEx := "界"
	chrEx := '界'
	fmt.Printf("Type of strEX  %T \n", strEx)
	fmt.Printf("Type of strEX[0]:  %T \n", strEx[0])
	fmt.Printf("Type of chrEx:  %T \n", chrEx)

}
