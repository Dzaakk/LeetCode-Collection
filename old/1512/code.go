func numIdenticalPairs(nums []int) int {
	ans := 0
	cnt := make(map[int]int)
	for _, v := range nums {
		ans += cnt[v]
		cnt[v]++
	}
	return ans
}