package containerwithmostwater

import "math"

func area(i int, j int, hi int, hj int) int {
	return (j - i) * int(math.Min(float64(hi), float64(hj)))
}

func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1
	for i < j {
		a := area(i, j, height[i], height[j])
		if a > max {
			max = a
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}
