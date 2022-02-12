package longestsubstring

import "fmt"

func example(s string) {
	fmt.Printf("Input: s = \"%v\"\n", s)
	fmt.Printf("Output: %v\n", lengthOfLongestSubstring(s))
}

func Examples() {
	fmt.Println("3. Longest Substring Without Repeating Characters")
	example("abcabcbb")
	example("bbbbb")
	example("pwwkew")
}
