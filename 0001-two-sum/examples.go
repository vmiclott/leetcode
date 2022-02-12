package twosum

import (
	"fmt"
)

func example(slice []int, target int) {
	fmt.Printf("Input: nums = %v, target = %v\n", slice, target)
	fmt.Printf("Output: %v\n", twoSum(slice, target))
}

func Examples() {
	fmt.Println("1. Two Sum")
	example([]int{2, 7, 11, 15}, 9)
	example([]int{3, 2, 4}, 6)
	example([]int{3, 3}, 6)
}
