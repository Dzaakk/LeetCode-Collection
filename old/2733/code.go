//02:27
func findNonMinOrMax(nums []int) int {
	sort.Ints(nums)

	if len(nums) < 3{
		return -1
	} else {
		return nums[len(nums)-2]
	}

}
