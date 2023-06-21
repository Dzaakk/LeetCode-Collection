func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if target == (nums[i] + nums[j]) {
				if i < j {
					nums = []int{i, j}
					fmt.Println("i < j")
				} else {
					nums = []int{j, i}
					fmt.Println("j < i")
				}
				break
			}
		}
	}

	return nums
}