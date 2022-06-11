package minopstoreducetozero

import "fmt"

func example(nums []int, x int) {
	fmt.Printf("Input: nums = %v, x = %v\n", nums, x)
	fmt.Printf("Output: %v\n", minOperations(nums, x))
}

func Examples() {
	fmt.Println("1658. Minimum Operations to Reduce X to Zero")
	example([]int{1, 1, 4, 2, 3}, 5)
	example([]int{5, 6, 7, 8, 9}, 4)
	example([]int{3, 2, 20, 1, 1, 3}, 10)
}
