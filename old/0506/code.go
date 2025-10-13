//09:16
func findRelativeRanks(score []int) []string {
	person := make(map[int]int)
	for i, s := range score {
		person[s]= i
	}

	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	result := make([]string, len(score))

	for i, s := range score {
		switch i {
		case 0:
			result[person[s]] = "Gold Medal"
		case 1:
			result[person[s]] = "Silver Medal"
		case 2:
			result[person[s]] = "Bronze Medal"
		default:
			result[person[s]] = strconv.Itoa(i + 1)
		}
	}
	return result
}