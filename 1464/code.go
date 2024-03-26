//02:42
func maxProduct(nums []int) int {
	sort.Ints(nums)

	result := (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)

	return result
}
func maxProduct(nums []int) int {
	max, secMax := 0, 0

	for _, num := range nums {
		if num > max {
			max, secMax = num, max
			continue
		}
		if num > secMax{
			secMax = num
		}
	}

	return (max-1) * (secMax-1)
}
