package romantointeger

func romanToInt(s string) int {
	var (
		i, res, n int = 0, 0, len(s)
		c, d      byte
	)

	for i < n {
		c = s[i]
		switch c {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			if i+1 < n {
				d = s[i+1]
				if d == 'D' {
					res += 400
					i++
				} else if d == 'M' {
					res += 900
					i++
				} else {
					res += 100
				}
			} else {
				res += 100
			}
		case 'L':
			res += 50
		case 'X':
			if i+1 < n {
				d = s[i+1]
				if d == 'L' {
					res += 40
					i++
				} else if d == 'C' {
					res += 90
					i++
				} else {
					res += 10
				}
			} else {
				res += 10
			}
		case 'V':
			res += 5
		case 'I':
			if i+1 < n {
				d = s[i+1]
				if d == 'V' {
					res += 4
					i++
				} else if d == 'X' {
					res += 9
					i++
				} else {
					res += 1
				}
			} else {
				res += 1
			}
		}
		i++
	}
	return res
}
