//03:00
func mostWordsFound(sentences []string) int {
	var res int

	for _, sentence := range sentences {
		count := 0
		for _, letter := range sentence {
			if letter == ' ' {
				count++
			}
		}
		if count > res {
			res = count
		}
	}
	return res+1
}
