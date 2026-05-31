package main

//15:10
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := range nums {
		result[i] = 1
	}

	prefix := 1
	for i := 0; i < len(nums); i++ {
		result[i] *= prefix
		prefix *= nums[i]
	}

	postfix := 1
	n := len(nums)
	for j := n - 1; j >= 0; j-- {
		result[j] *= postfix
		postfix *= nums[j]
	}

	return result

}
