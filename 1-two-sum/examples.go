package twosum

import (
	"fmt"
)

func Examples() {
	fmt.Println("1. Two Sum")
	fmt.Println("Input: nums = [2,7,11,15], target = 9")
	fmt.Printf("Output: %v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Input: nums = [3,2,4], target = 6")
	fmt.Printf("Output: %v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Println("Input: nums = [3,3], target = 6")
	fmt.Printf("Output: %v\n", twoSum([]int{3, 3}, 6))
}
