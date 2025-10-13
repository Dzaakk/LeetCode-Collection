func tictactoe(moves [][]int) string {
	grid := [3][3]byte{}

	for i, move := range moves {
		row, col := move[0], move[1]
		if i%2 == 0 {
			grid[row][col] = 'X'
		} else {
			grid[row][col] = 'O'
		}

		if checkWin(grid, row, col) {
			if i%2 == 0 {
				return "A"
			} else {
				return "B"
			}
		}
	}

	if len(moves) == 9 {
		return "Draw"
	} else {
		return "Pending"
	}
}

func checkWin(grid [3][3]byte, row, col int) bool {
	player := grid[row][col]
	fmt.Println("Player =", player)

	// Check for a winning condition in the row
	if grid[row][0] == player && grid[row][1] == player && grid[row][2] == player {
		return true
	}

	// Check for a winning condition in the column
	if grid[0][col] == player && grid[1][col] == player && grid[2][col] == player {
		return true
	}

	// Check for a winning condition in the diagonal (top-left to bottom-right)
	if row == col && grid[0][0] == player && grid[1][1] == player && grid[2][2] == player {
		return true
	}

	// Check for a winning condition in the diagonal (top-right to bottom-left)
	if row+col == 2 && grid[0][2] == player && grid[1][1] == player && grid[2][0] == player {
		return true
	}

	return false
}