package integertoroman

func intToRoman(num int) string {
	var res string

	for num >= 1000 {
		res += "M"
		num -= 1000
	}
	if num >= 900 {
		res += "CM"
		num -= 900
	}
	for num >= 500 {
		res += "D"
		num -= 500
	}
	if num >= 400 {
		res += "CD"
		num -= 400
	}
	for num >= 100 {
		res += "C"
		num -= 100
	}
	if num >= 90 {
		res += "XC"
		num -= 90
	}
	if num >= 50 {
		res += "L"
		num -= 50
	}
	if num >= 40 {
		res += "XL"
		num -= 40
	}
	for num >= 10 {
		res += "X"
		num -= 10
	}
	if num == 9 {
		return res + "IX"
	}
	if num >= 5 {
		res += "V"
		num -= 5
	}
	if num == 4 {
		return res + "IV"
	}
	for num >= 1 {
		res += "I"
		num -= 1
	}
	return res
}
