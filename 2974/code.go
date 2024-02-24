//06:02
func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i+= 2{
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}
func numberGame(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i+1], nums[i])
	}
	return res
}