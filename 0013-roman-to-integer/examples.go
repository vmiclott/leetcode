package romantointeger

import "fmt"

func example(s string) {
	fmt.Printf("Input: s = \"%v\"\n", s)
	fmt.Printf("Output: %v\n", romanToInt(s))
}

func Examples() {
	fmt.Println("13. Roman to Integer")
	example("III")
	example("LVIII")
	example("MCMXCIV")
}
