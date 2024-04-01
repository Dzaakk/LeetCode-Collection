//16:00
func sortArrayByParity(nums []int) []int {
	sort.Ints(nums)

	result := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for _, num := range nums {
		if num % 2 == 0 {
			result[left] = num
			left++
		} else {
			result[right] = num
			right--
		}
	}
	return result
}
