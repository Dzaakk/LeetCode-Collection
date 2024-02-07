func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, 0)
	biggest := 0
	for _, candy := range candies {
		if candy > biggest {
			biggest = candy
		}
	}
	for _, candy := range candies {
		if candy+extraCandies >= biggest {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
//re commit wrong date