package lettercombinations

import "fmt"

func example(digits string) {
	fmt.Printf("Input: digits = \"%v\"\n", digits)
	fmt.Printf("Output: %v\n", letterCombinations(digits))
}

func Examples() {
	fmt.Println("17. Letter Combinations of a Phone Number")
	example("23")
	example("")
	example("2")
}
