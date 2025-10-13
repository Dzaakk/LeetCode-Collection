func maxDistance(position []int, m int) int {
	sort.Ints(position)
	l, r := 1, position[len(position)-1]-position[0]
	ans := -1

	for l <= r {
		mid := l + (r-l)/2
		lastPosition, balls := position[0], 1

		for i := 1; i < len(position); i++ {
			if position[i]-lastPosition >= mid {
				lastPosition = position[i]
				balls++
			}
		}
		if balls >= m {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}