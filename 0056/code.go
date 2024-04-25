//12:31
func merge(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result = append(result, intervals[0])
	for _, v:= range intervals{
		last := result[len(result)-1]
		if v[0] <= last[1] {
			if v[1] > last[1] {
				last[1] = v[1]
			}
		} else {
			result = append(result, v)
		}
	}

	return result
}