package longeststringchain

import "fmt"

func example(words []string) {
	fmt.Printf("Input: words = %v\n", words)
	fmt.Printf("Output: %v\n", longestStrChain(words))
}

func Examples() {
	fmt.Println("1048. Longest String Chain")
	example([]string{"a", "b", "ba", "bca", "bda", "bdca"})
	example([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"})
	example([]string{"abcd", "dbqca"})
}
