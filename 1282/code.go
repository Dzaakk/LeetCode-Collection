//51:03
func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)
	result := [][]int{}

	for i, size := range groupSizes {
		if _, exist := groups[size]; !exist{
			groups[size] = []int{}
		}

		groups[size] = append(groups[size], i)
		if len(groups[size]) == size{
			result= append(result, groups[size])
			delete(groups, size)
		}
	}
	return result
}
