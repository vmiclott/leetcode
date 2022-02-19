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

func equalAtPos(strs []string, pos int) bool {
	c := strs[0][pos]
	for _, s := range strs {
		if c != s[pos] {
			return false
		}
	}
	return true
}

func longestCommonPrefix(strs []string) string {
	l, n := 0, minLength(strs)
	for i := 0; i < n; i++ {
		if equalAtPos(strs, i) {
			l++
		} else {
			return strs[0][0:l]
		}
	}
	return strs[0][0:l]
}
