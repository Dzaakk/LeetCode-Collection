//03:08
func maxProductDifference(nums []int) int {
	sort.Ints(nums)

	left := nums[0] * nums [1]
	right := nums[len(nums)-1] * nums[len(nums)-2]

	result := right - left

	return result
}
