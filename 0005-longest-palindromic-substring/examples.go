package longestpalindromicsubstring

import "fmt"

func example(s string) {
	fmt.Printf("Input: s = \"%v\"\n", s)
	fmt.Printf("Output: \"%v\"\n", longestPalindrome(s))
}

func Examples() {
	fmt.Println("5. Longest Palindromic Substring")
	example("babad")
	example("cbbd")
}
