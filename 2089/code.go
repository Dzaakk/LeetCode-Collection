func targetIndices(nums []int, target int) []int {
	out := make([]int, 0)
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	for i, v := range nums {
		if v == target {
			out = append(out, i)
		}
	}
	return out
}

//10.16

func targetIndices(nums []int, target int) []int {
	var out []int
	count, left := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}

		if nums[i] < target {
			left++
		}
	}

	for i := 0; i < count; i++ {
		out = append(out, left)
		left++
	}

	return out
}