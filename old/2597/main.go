func beautifulSubsets(nums []int, k int) int {
	count := 0
	n := len(nums)
	visited := make(map[int]int)

	var explore func(index int)
	explore = func(index int) {
		if index == n {
			count++
			return
		}

		num := nums[index]

		if _, exist1 := visited[num-k]; !exist1 {
			if _, exist2 := visited[num+k]; !exist2 {
				visited[num]++
				explore(index + 1)
				visited[num]--
				if visited[num] == 0 {
					delete(visited, num)
				}
			}
		}
		explore(index + 1)
	}
	explore(0)

	return count - 1
}