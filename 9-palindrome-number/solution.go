package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	digits := []int{}
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	i, j := 0, len(digits)-1
	for i < j {
		if digits[i] != digits[j] {
			return false
		}
		i++
		j--
	}

	return true
}
