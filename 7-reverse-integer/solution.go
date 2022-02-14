package reverseinteger

import "math"

func reverse(x int) int {
	var (
		sign, num, d, res, r int32
	)

	if x == math.MinInt32 || x == math.MaxInt32 {
		return 0
	}

	sign = 1
	if x < 0 {
		sign = -1
	}
	num = int32(x) * sign

	if num < 10 {
		return x
	}

	d = 10
	for num/d > 9 {
		d *= 10
	}

	res = 0
	for num > 0 {
		r = num % 10
		num = (num - r) / 10
		res += r * d
		if res/d%10 != r {
			return 0
		}
		d /= 10
	}

	return int(sign * res)
}
