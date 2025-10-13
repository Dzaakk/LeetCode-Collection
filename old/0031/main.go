func nextPermutation(nums []int) {
	j := len(nums) - 1
	for ; j > 0 && nums[j-1] >= nums[j]; j-- {

	}

	if j != 0 {
		l := len(nums) - 1
		for ; l > j-1 && nums[j-1] > nums[l]; l-- {
		}

		nums[j-1], nums[l] = nums[l], nums[j-1]
	}

	for k := len(nums) - 1; j < k; j, k = j+1, k-1 {
		nums[j], nums[k] = nums[k], nums[j]

	}
}