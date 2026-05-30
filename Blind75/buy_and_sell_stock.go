package main

//18:25
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	profit := 0
	l, r := 0, 1
	for r < len(prices) {
		if prices[l] < prices[r] {
			x := prices[r] - prices[l]
			if x > profit {
				profit = x
			}
		}
		if prices[l] > prices[r] {
			l++
			continue
		}
		r++
	}

	return profit
}
