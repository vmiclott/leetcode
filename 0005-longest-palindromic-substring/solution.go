package longestpalindromicsubstring

func longestPalindrome(s string) string {
	n := len(s)
	start := 0
	end := 0
	longest := s[start : end+1]
	maxLength := end + 1 - start

	for i := range s {
		left := i
		right := i

		for right+1 < n && s[i] == s[right+1] {
			right += 1
		}

		for left-1 > 0 && s[left-1] == s[i] {
			left -= 1
		}

		for left-1 >= 0 && right+1 < n && s[left-1] == s[right+1] {
			left -= 1
			right += 1
		}

		length := right + 1 - left

		if length > maxLength {
			maxLength = length
			longest = s[left : right+1]
		}
	}
	return longest
}
