package mediantwosortedarrays

import "math"

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func findMedianSortedArray(nums []int) float64 {
	length := len(nums)
	if length%2 == 0 {
		return float64(nums[length/2-1])/2 + float64(nums[length/2])/2
	}
	return float64(nums[length/2])
}

func findMedianSortedArraysLongShort(long []int, short []int) float64 {
	lLong, lShort := len(long), len(short)
	lMerged := lShort + lLong
	isEven := lMerged%2 == 0
	halfLMerged := lMerged / 2

	if lShort == 0 {
		if isEven {
			return float64(long[halfLMerged-1]+long[halfLMerged]) / 2
		}
		return float64(long[halfLMerged-1])
	}

	i, j := int(math.Max(-1, float64(halfLMerged)-float64(lShort)-1)), halfLMerged-1 // binary search index boundaries for long slice
	for i <= j {
		var (
			m, n, longMax, shortMax, longNext, shortNext int
		)
		m = i + (j-i)/2         // right index of partition of long slice
		n = halfLMerged - m - 2 // right index of partition of short slice

		if m < 0 {
			longMax = minInt // no elements in left partition of long slice
		} else {
			longMax = long[m]
		}

		if n < 0 {
			shortMax = minInt // no elements in left partition of short slice
		} else {
			shortMax = short[n]
		}

		if lLong < m+2 {
			longNext = maxInt // no elements in long slice past longMax
		} else {
			longNext = long[m+1]
		}

		if lShort < n+2 {
			shortNext = maxInt // no elements in short slice past shortMax
		} else {
			shortNext = short[n+1]
		}

		if longMax <= shortNext && shortMax <= longNext {
			if isEven {
				return (math.Max(float64(longMax), float64(shortMax)) + math.Min(float64(longNext), float64(shortNext))) / 2
			}
			return math.Min(float64(longNext), float64(shortNext))
		}

		if longMax > shortNext {
			j = m - 1
		} else {
			// shortMax > longNext
			i = m + 1
		}
	}

	return 0
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return findMedianSortedArray(nums2)
	}
	if len(nums2) == 0 {
		return findMedianSortedArray(nums1)
	}

	if len(nums1) > len(nums2) {
		return findMedianSortedArraysLongShort(nums1, nums2)
	}
	return findMedianSortedArraysLongShort(nums2, nums1)
}
