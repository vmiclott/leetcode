package palindromenumber

import "fmt"

func example(x int) {
	fmt.Printf("Input: x = %v\n", x)
	fmt.Printf("Output: %v\n", isPalindrome(x))
}

func Examples() {
	fmt.Println("9. Palindrome Number")
	example(121)
	example(-121)
	example(10)
}
