package mediantwosortedarrays

func mergeSortedArrays(nums1 []int, nums2 []int) []int {
	merged := []int{}
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i >= len(nums1) {
			merged = append(merged, nums2[j:]...)
			break
		} else if j >= len(nums2) {
			merged = append(merged, nums1[i:]...)
			break
		} else if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	return merged
}

func findMedianSortedArray(nums []int) float64 {
	length := len(nums)
	if length%2 == 0 {
		return float64(nums[length/2-1])/2 + float64(nums[length/2])/2
	}
	return float64(nums[length/2])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return findMedianSortedArray(nums2)
	}
	if len(nums2) == 0 {
		return findMedianSortedArray(nums1)
	}

	return findMedianSortedArray(mergeSortedArrays(nums1, nums2))
}
