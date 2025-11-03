package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	l, r := 0, 1

	maxP := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			tmp := prices[r] - prices[l]
			if tmp > maxP {
				maxP = tmp
			}
		}
		if prices[l] > prices[r] {
			l++
			continue
		}
		r++
	}
	return maxP
}

// func maxProfit(prices []int) int {
// 	if len(prices) == 1 {
// 		return 0
// 	}

// 	minP := prices[0]
// 	maxP := prices[0]

// 	maxProfit := 0

// 	for _, price := range prices {
// 		if price > maxP {
// 			maxProfit = max(maxProfit, price-minP)
// 			maxProfit = price
// 		} else if price < minP {
// 			minP = price
// 			maxP = price
// 		}
// 	}

// 	return maxProfit
// }
