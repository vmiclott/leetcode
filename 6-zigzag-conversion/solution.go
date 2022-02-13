package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	lines := make([]string, numRows)
	m := 2*numRows - 2

	for i, r := range s {
		im := i % m
		if im < numRows {
			lines[im] = lines[im] + string(r)
		} else {
			lines[2*numRows-2-im] = lines[2*numRows-2-im] + string(r)
		}
	}

	return strings.Join(lines, "")
}
