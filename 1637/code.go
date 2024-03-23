package _637
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	maxWidth := 0
	for i := 1; i < len(points); i++ {
		maxWidth = max(maxWidth, points[i][0] - points[i-1][0])
	}

	return maxWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//solution 2
func maxWidthOfVerticalArea(points [][]int) int {
	exes := make([]int, len(points))
	for i := range points {
		exes[i] = points[i][0]
	}

	sort.Ints(exes)
	max := 0
	for i := 0; i < len(exes) - 1; i++ {
		if max < exes[i+1] - exes[i] {
			max = exes[i+1] - exes[i]
		}
	}

	return max
}
