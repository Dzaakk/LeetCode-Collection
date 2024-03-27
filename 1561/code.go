//14:20
func maxCoins(piles []int) int {
	sort.Ints(piles)
	var counter int
	for i := len(piles)/3 ; i < len(piles); i+=2 {
		counter += piles[i]
	}

	return counter

}
