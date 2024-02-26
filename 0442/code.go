//25:37
func findDuplicates(nums []int) []int {
	countMap := make(map[int]int)

	duplicates := []int{}

	for _, v := range nums {
		if _, exists := countMap[v]; exists {
			duplicates = append(duplicates, v)
		} else {
			countMap[v] = 1
		}
	}

	return duplicates
}
