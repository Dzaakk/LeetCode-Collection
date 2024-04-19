//05:10
func sumOfUnique(nums []int) int {
	mapNum := make(map[int]int)
	answer := 0
	for _, num := range nums {
		mapNum[num]++
	}

	for i := 0; i < len(nums); i++ {
		if count := mapNum[nums[i]]; count == 1{
			answer += nums[i]
		}
	}

	return answer
}
