package twosumii

import "fmt"

func example(numbers []int, target int) {
	fmt.Printf("Input: numbers = %v, target = %v\n", numbers, target)
	fmt.Printf("Output: %v\n", twoSum(numbers, target))
}

func Examples() {
	fmt.Println("167. Two Sum II - Input Array Is Sorted")
	example([]int{2, 7, 11, 15}, 9)
	example([]int{2, 3, 4}, 6)
	example([]int{-1, 0}, -1)
}
