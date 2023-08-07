func checkStraightLine(coordinates [][]int) bool {
	n := len(coordinates)

	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]

	for i := 1; i < n; i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if (y-y1)*(x1-x0) != (y1-y0)*(x-x1) {
			return false
		}
	}
	return true
}