package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = 10*reverted + x%10
		x /= 10
	}

	return x == reverted || x == reverted/10
}
