func findWordsContaining(words []string, x byte) []int {
	result := []int{}
	newX := int32(x)
	for i, word := range words {
		for _, w := range word {
			if newX == w {
				result = append(result, i)
				break
			}
			continue
		}
	}
	return result
}
