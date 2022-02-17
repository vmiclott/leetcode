package regularexpressionmatching

import "fmt"

func example(s string, p string) {
	fmt.Printf("Input: s = \"%v\", p = \"%v\"\n", s, p)
	fmt.Printf("Output: %v\n", isMatch(s, p))
}

func Examples() {
	fmt.Println("10. Regular Expression Matching")
	example("aa", "a")
	example("aa", "a*")
	example("ab", ".*")
}
