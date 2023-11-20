func countNegatives(grid [][]int) int {
	res := 0
	for _, v := range grid {
		for _, x := range v {
			if x < 0 {
				res++
			}
		}
	}
	return res
}

// 02.00

func countNegatives(grid [][]int) int {
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}
	return count
}