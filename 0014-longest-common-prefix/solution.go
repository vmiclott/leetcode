package longestcommonprefix

func minLength(strs []string) int {
	res := len(strs[0])
	for _, s := range strs {
		if len(s) < res {
			res = len(s)
		}
	}
	return res
}

func equalPrefixes(strs []string, l int) bool {
	prefix := strs[0][0:l]
	for _, s := range strs {
		if prefix != s[0:l] {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	start, end := 1, minLength(strs)
	if end == 0 {
		return ""
	}
	for start <= end {
		m := (start + end) / 2
		if equalPrefixes(strs, m) {
			start = m + 1
		} else {
			end = m - 1
		}
	}
	return strs[0][0 : (start+end)/2]
}
