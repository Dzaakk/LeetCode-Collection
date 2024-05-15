func findingUsersActiveMinutes(logs [][]int, k int) []int {
	output := make([]int, k)
	userActiveMinute := make(map[int]map[int]struct{})
	n := len(logs)

	for i := 0; i < n; i++ {
		user, minute := logs[i][0], logs[i][1]

		if _, ok := userActiveMinute[user]; !ok {
			userActiveMinute[user] = map[int]struct{}{}
		}
		userActiveMinute[user][minute] = struct{}{}
	}

	for _, v := range userActiveMinute {
		if len(v)-1 >= k {
			continue
		}

		output[len(v)-1]++
	}

	return output
}