package containerwithmostwater

import "fmt"

func example(height []int) {
	fmt.Printf("Input: height = %v\n", height)
	fmt.Printf("Output: %v\n", maxArea(height))
}

func Examples() {
	fmt.Println("11. Container With Most Water")
	example([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	example([]int{1, 1})
}
