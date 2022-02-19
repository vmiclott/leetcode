package longestcommonprefix

import "fmt"

func example(strs []string) {
	fmt.Printf("Input: strs = %v\n", strs)
	fmt.Printf("Output: \"%v\"\n", longestCommonPrefix(strs))
}

func Examples() {
	fmt.Println("14. Longest Common Prefix")
	example([]string{"flower", "flow", "flight"})
	example([]string{"dog", "racecar", "car"})
}
