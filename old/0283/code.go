func moveZeroes(nums []int) {
	i := 0

	for v := range nums {
		if nums[v] != 0 {
			nums[i] = nums[v]
			i++
		}
	}

	for ; i < len(nums); i++ {
		nums[i] = 0
	}

	return

}