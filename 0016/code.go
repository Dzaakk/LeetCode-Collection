//06:38
func threeSumClosest(nums []int, target int) int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	sort.Ints(nums)

	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == target {
				return sum
			}

			if abs(target-sum) < abs(target-result) {
				result = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
