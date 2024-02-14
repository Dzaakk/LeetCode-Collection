//54:42
func threeSum(nums []int) [][]int {
	res := make([][]int, 0, len(nums)/3)
	sort.Ints(nums)
	var sum, left, right, tempLeft int

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		left, right = i + 1, len(nums)- 1

		for left < right {
			if nums[left] == tempLeft {
				left++
			}
			sum = num + nums[left] + nums[right]

			if sum == 0 {
				res = append(res, []int{num, nums[left], nums[right]})
				left++

				for left < right && nums[left] == nums[left -1]{
					left++
				}
			}
			tempLeft = nums[left]
			if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
