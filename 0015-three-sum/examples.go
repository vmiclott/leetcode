package threesum

import "fmt"

func example(nums []int) {
	fmt.Printf("Input: nums = %v\n", nums)
	fmt.Printf("Output: %v\n", threeSum(nums))
}

func Examples() {
	fmt.Println("15. 3Sum")
	example([]int{-1, 0, 1, 2, -1, -4})
	example([]int{})
	example([]int{0})
}
