func finalValueAfterOperations(operations []string) int {
	var result int
	for _, v := range operations {
		if v == "--X" || v == "X--" {
			result--
		} else if v == "++X" || v == "X++" {
			result++
		}
	}
	return result
}