package threesumclosest

import "fmt"

func example(nums []int, target int) {
	fmt.Printf("Input: nums = %v, target = %v\n", nums, target)
	fmt.Printf("Output: %v\n", threeSumClosest(nums, target))
}

func Examples() {
	fmt.Println("16. 3Sum Closest")
	example([]int{-1, 2, 1, -4}, 1)
	example([]int{0, 0, 0}, 1)
}
