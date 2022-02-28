package lettercombinations

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	n := 1
	for _, r := range digits {
		if r == '7' || r == '9' {
			n *= 4
		} else {
			n *= 3
		}
	}

	res := make([]string, n)
	m := n
	for _, r := range digits {
		switch r {
		case '2':
			m /= 3
			chars := []string{"a", "b", "c"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%3]
			}
		case '3':
			m /= 3
			chars := []string{"d", "e", "f"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%3]
			}
		case '4':
			m /= 3
			chars := []string{"g", "h", "i"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%3]
			}
		case '5':
			m /= 3
			chars := []string{"j", "k", "l"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%3]
			}
		case '6':
			m /= 3
			chars := []string{"m", "n", "o"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%3]
			}
		case '7':
			m /= 4
			chars := []string{"p", "q", "r", "s"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%4]
			}
		case '8':
			m /= 3
			chars := []string{"t", "u", "v"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%3]
			}
		case '9':
			m /= 4
			chars := []string{"w", "x", "y", "z"}
			for i := 0; i < n; i++ {
				res[i] += chars[(i/m)%4]
			}
		}
	}
	return res
}
