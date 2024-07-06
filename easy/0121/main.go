func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if (prices[i] - min) > profit {
			profit = prices[i] - min
		}
	}

	return profit
}