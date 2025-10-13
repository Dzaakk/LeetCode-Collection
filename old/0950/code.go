//29:12
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	result := make([]int, len(deck))
	for head := len(deck); head > 0; {
		if len(deck)-head > 0 {
			tmp := result[len(result)-1]
			for i := len(result) - 1; i > head; i--{
				result[i] = result[i-1]
			}
			result[head] = tmp
		}
		head--
		result[head] = deck[head]
	}

	return result
}
