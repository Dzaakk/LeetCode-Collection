func numberOfSubarrays(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		nums[i] %= 2
	}

	prefixCount := make([]int, len(nums)+1)
	prefixCount[0] = 1
	s := 0
	res := 0

	for _, num := range nums {
		s += num
		if s >= k {
			res += prefixCount[s-k]
		}
		prefixCount[s]++
	}

	return res
}