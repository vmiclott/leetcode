package maximumerasurevalue

func maximumUniqueSubarray(nums []int) int {
	start, sum, maxSum := 0, 0, 0

	firstEl := nums[0]
	lastIdx := make(map[int]int)

	for end, endEl := range nums {
		if start == 0 && end > 0 && firstEl == endEl {
			if sum > maxSum {
				maxSum = sum
			}
			sum -= firstEl
			start++
		} else if idx := lastIdx[endEl]; idx > 0 && idx >= start {
			if sum > maxSum {
				maxSum = sum
			}
			for i := start; i <= idx; i++ {
				sum -= nums[i]
			}
			start = idx + 1
		}

		sum += endEl
		lastIdx[endEl] = end
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
