package longestsubstring

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	length, maxLength, start := 0, 0, 0
	runeMap := map[rune]int{}
	for i, r := range s {
		if j, ok := runeMap[r]; ok {
			for k := start; k <= j; k++ {
				delete(runeMap, rune(s[k]))
				length -= 1
			}
			start = j + 1
		}

		length += 1
		if length > maxLength {
			maxLength = length
		}
		runeMap[r] = i
	}
	if maxLength == 0 {
		return len(s)
	}
	return maxLength
}
