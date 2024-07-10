//05:35
func maxSubArray(nums []int) int {
	current := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		current = max(nums[i]+current, nums[i])

		result = max(result, current)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
