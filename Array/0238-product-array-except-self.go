package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := range nums {
		res[i] = 1
	}

	prefix := 1
	for i := 0; i < len(nums); i++ {
		res[i] *= prefix
		prefix *= nums[i]
	}

	postfix := 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] *= postfix
		postfix *= nums[j]
	}

	return res
}
