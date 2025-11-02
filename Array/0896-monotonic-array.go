package main

func isMonotonic(nums []int) bool {
	increase, decrease := 0, 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			increase++
		}
		if nums[i] > nums[i+1] {
			decrease++
		}
	}
	if increase != 0 && decrease != 0 {
		return false
	}

	return true
}
