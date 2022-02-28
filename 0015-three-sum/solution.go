package threesum

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 1 {
		return [][]int{}
	}
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		left := i + 1
		right := n - 1
		val := nums[i]

		for left < right {
			sum := val + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{val, nums[left], nums[right]})
				left++
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else if sum < 0 {
				left++
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else if sum > 0 {
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			}
		}
	}

	return result
}
