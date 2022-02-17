package regularexpressionmatching

func isMatch(s string, p string) bool {
	memo := map[int]map[int]bool{}

	var d func(i int, j int) bool
	d = func(i int, j int) bool {
		if _, ok := memo[i]; !ok {
			memo[i] = map[int]bool{}
		}
		if _, ok := memo[i][j]; ok {
			return memo[i][j]
		}
		if j == len(p) {
			memo[i][j] = i == len(s)
			return memo[i][j]
		}

		match := i < len(s) && (p[j] == s[i] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			memo[i][j] = d(i, j+2) || (match && d(i+1, j))
			return memo[i][j]
		}

		memo[i][j] = match && d(i+1, j+1)
		return memo[i][j]
	}

	return d(0, 0)
}
