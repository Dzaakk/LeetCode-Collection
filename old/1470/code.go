func shuffle(nums []int, n int) []int {
	x := []int{}
	y := []int{}
	result := []int{}
	for _, num := range nums[:n] {
		x = append(x, num)
	}

	for _, num := range nums[n:] {
		y = append(y, num)
	}

	for i := 0; i < n; i++ {
		result = append(result, x[i], y[i])
	}

	return result
}