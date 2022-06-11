package minopstoreducetozero

func minOperations(nums []int, x int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	if sum < x {
		return -1
	}

	if sum == x {
		return len(nums)
	}

	maxLength, start, curr, target := 0, 0, 0, sum-x

	for end := range nums {
		curr += nums[end]

		for start <= end && curr > target {
			curr -= nums[start]
			start++
		}

		if curr == target {
			currLength := end - start + 1
			if currLength > maxLength {
				maxLength = currLength
			}
		}
	}

	if maxLength == 0 && sum != x {
		return -1
	}

	return len(nums) - maxLength
}
