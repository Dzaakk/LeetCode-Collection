//09:09
func missingNumber(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return ((len(nums)+1)*len(nums))/2 - sum
}
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if num != i {
			return i
		}
	}

	return len(nums)
}
