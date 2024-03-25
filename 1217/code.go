//05:02
func minCostToMoveChips(position []int) int {
	evens, odds := 0, 0
	for _, pos := range position {
		if pos%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	if odds < evens {
		return odds
	}

	return evens
}
