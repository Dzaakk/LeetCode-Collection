//22:59
func deleteGreatestValue(grid [][]int) int {
	for _, g := range grid {
		sort.Ints(g)
	}

	var result int
	for i := 0; i < len(grid[0]); i++ {
		tmp := 0
		for j := 0; j < len(grid); j++ {
			tmp = max(tmp, grid[j][i])
		}
		result += tmp
	}
	return result
}

