package longestsubstring

import "fmt"

func Examples() {
	fmt.Println("Input: s = \"abcabcbb\"")
	fmt.Printf("Output: %v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Println("Input: s = \"bbbbb\"")
	fmt.Printf("Output: %v\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Println("Input: s = \"pwwkew\"")
	fmt.Printf("Output: %v\n", lengthOfLongestSubstring("pwwkew"))
}
