package integertoroman

import "fmt"

func example(num int) {
	fmt.Printf("Input: num = %v\n", num)
	fmt.Printf("Output: %v\n", intToRoman(num))
}

func Examples() {
	fmt.Println("12. Integer to Roman")
	example(3)
	example(58)
	example(1994)
}
