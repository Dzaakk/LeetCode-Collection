package _448
//23:37
func findDisappearedNumbers(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}

	result := make([]int, 0)
	for _, n := range nums {
		result[n-1] = n
	}

	j := 0
	for i := 0; i < length; i++ {
		if result[i] == 0 {
			result[j] = i + 1
			j++
		}
	}
	return result[:j]
}
//sec
func findDisappearedNumbers(nums []int) []int {
	for i, _ := range nums {
		for nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	result := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1{
			result = append(result, i+1)
		}
	}
	return result
}