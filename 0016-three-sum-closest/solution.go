package threesumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		val := nums[i]
		left := i + 1
		right := n - 1

		for left < right {
			sum := val + nums[left] + nums[right]
			if sum == target {
				return target
			}
			if sum < target {
				left++
				for nums[left-1] == nums[left] && left < right {
					left++
				}
			} else if sum > target {
				right--
				for nums[right] == nums[right+1] && left < right {
					right--
				}
			}
			if int(math.Abs(float64(sum)-float64(target))) < int(math.Abs(float64(result)-float64(target))) {
				result = sum
			}
		}
	}
	return result
}
