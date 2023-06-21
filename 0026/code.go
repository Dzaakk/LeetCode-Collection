func removeDuplicates(nums []int) int {
	result := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[i-1] = nums[i]
		}
	}
	return result + 1
}