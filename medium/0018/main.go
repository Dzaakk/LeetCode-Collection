func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
		for j := i+1; j < len(nums); j++ {
			for j > i+1 && j < len(nums) && nums[j] == nums[j-1] {
				j++
			}
			start := j+1
			end := len(nums)-1

			for start < end {
				if (nums[start]+nums[end]+nums[i]+nums[j]) < target {
					start++
					for start < len(nums) && nums[start] == nums[start-1] {
						start++
					}
				} else if (nums[start]+nums[end]+nums[i]+nums[j]) > target {
					end--
					for end > 0 && nums[end] == nums[end+1] {
						end--
					}
				} else {
					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					end--
					for start < len(nums) && nums[start] == nums[start-1] {
						start++
					}
					for end > 0 && nums[end] == nums[end+1] {
						end--
					}
				}
			}
		}
	}
	return result
}
