package stringtointeger

import "fmt"

func example(s string) {
	fmt.Printf("Input: s = \"%v\"\n", s)
	fmt.Printf("Output: %v\n", myAtoi(s))
}

func Examples() {
	fmt.Println("8. String to Integer (atoi)")
	example("42")
	example("   -42")
	example("4193 with words")
}
