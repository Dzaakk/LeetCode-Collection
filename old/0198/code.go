func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max1 := nums[0]
	max2 := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		max2, max1 = max(max1+nums[i], max2), max2
	}

	return max2

}