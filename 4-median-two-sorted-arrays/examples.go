package mediantwosortedarrays

import (
	"fmt"
)

func Examples() {
	fmt.Println("Input: nums1 = [1,3], nums2 = [2]")
	fmt.Printf("Output: %v\n", findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println("Input: nums1 = [1,2], nums2 = [3,4]")
	fmt.Printf("Output: %v\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
