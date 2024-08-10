func searchInsert(nums []int, target int) int {
	out := 0
	for i := range nums {
		if nums[i] == target {
			out = i
			break
		}
		if nums[i] < target {
			out += 1
		}
		if target > nums[len(nums)-1] {
			out = len(nums)
		}
	}
	return out
}