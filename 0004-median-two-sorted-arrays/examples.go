package mediantwosortedarrays

import (
	"fmt"
)

func example(nums1 []int, nums2 []int) {
	fmt.Printf("Input: nums1 = %v, nums2 = %v\n", nums1, nums2)
	fmt.Printf("Output: %v\n", findMedianSortedArrays(nums1, nums2))
}

func Examples() {
	fmt.Println("4. Median of Two Sorted Arrays")
	example([]int{1, 3}, []int{2})
	example([]int{1, 2}, []int{3, 4})
}
