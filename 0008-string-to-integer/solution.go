package stringtointeger

import "math"

var digitsMap = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func isDigit(c byte) bool {
	_, ok := digitsMap[c]
	return ok
}

func myAtoi(s string) int {
	i := 0
	n := len(s)
	sign := 1
	res := 0

	for i < n && s[i] == ' ' {
		i++
	}

	if i < n && s[i] == '-' {
		sign = -1
		i++
	} else if i < n && s[i] == '+' {
		i++
	}

	for i < n && s[i] == '0' {
		i++
	}

	for i < n && isDigit(s[i]) {
		res *= 10
		res += digitsMap[s[i]]

		if sign*res < int(math.MinInt32) {
			return int(math.MinInt32)
		}

		if sign*res > int(math.MaxInt32) {
			return int(math.MaxInt32)
		}

		i++
	}

	return sign * res
}
