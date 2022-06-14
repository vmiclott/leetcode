package maximumerasurevalue

import "fmt"

func example(nums []int) {
	fmt.Printf("Input: nums = %v\n", nums)
	fmt.Printf("Output: %v\n", maximumUniqueSubarray(nums))
}

func Examples() {
	example([]int{4, 2, 4, 5, 6})
	example([]int{5, 2, 1, 2, 5, 2, 1, 2, 5})
}
