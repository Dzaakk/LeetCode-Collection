func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	counter := make(map[int]int)

	for _, num := range nums {
		counter[num]++
	}

	var backtrack func(comb []int, n int)

	backtrack = func(comb []int, n int) {

		if len(comb) == n {
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			result = append(result, tmp)
			return
		}

		for num, count := range counter {
			if count == 0 {
				continue
			}

			comb = append(comb, num)
			counter[num]--
			backtrack(comb, n)

			counter[num]++
			comb = comb[:len(comb)-1]
		}
	}

	backtrack([]int{}, len(nums))

	return result
}