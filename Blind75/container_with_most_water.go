package main

//04:56
func maxArea(height []int) int {
	res := 0
	length := len(height)
	l, r := 0, length-1

	for l < r {
		area := min(height[l], height[r]) * (length - 1)

		if area > res {
			res = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
		length--
	}

	return res
}
