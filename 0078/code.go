//07:35
func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})

	for i := 0; i < len(nums); i++ {
		length := len(result)

		for j := 0; j < length; j++ {
			item := result[j]

			tmp := make([]int, len(item))
			copy(tmp, item)
			tmp = append(tmp, nums[i])

			result = append(result, tmp)
		}
	}
	return result
}

