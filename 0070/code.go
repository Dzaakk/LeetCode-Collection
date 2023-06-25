func climbStairs(n int) int {
	result, firstLast, secondLast := 0, 0, 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			result = 1
		} else if i == 2 {
			result = 2
		} else {
			result = firstLast + secondLast
		}

		secondLast = firstLast
		firstLast = result
	}
	return result
}