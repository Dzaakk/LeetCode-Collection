//08:44
func removeDuplicates(nums []int) int {
	mapNums := make(map[int]int)
	count := 0

	for i := 0; i < len(nums); i++ {
		if mapNums[nums[i]] < 2 {
			mapNums[nums[i]]++
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
func removeDuplicates(nums []int) int {
	var count int
	for _, num := range nums {
		if count == 0 || count == 1 || nums[count-2] != num {
			nums[count]= num
			count++
		}
	}
	return count
}