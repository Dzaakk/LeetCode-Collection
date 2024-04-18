//26:19
func findDifference(nums1 []int, nums2 []int) [][]int {
	mapNum1 := make(map[int]bool)
	mapNum2 := make(map[int]bool)
	answer := make([][]int, 2)

	for _, n:= range nums1 {
		mapNum1[n] = true
	}

	for _, n:= range nums2 {
		mapNum2[n] = true
	}

	for num, _ := range mapNum1 {
		if _, ok := mapNum2[num]; !ok {
			answer[0] = append(answer[0], num)
		}
	}
	for num, _ := range mapNum2 {
		if _, ok := mapNum1[num]; !ok {
			answer[1] = append(answer[1], num)
		}
	}

	return answer
}

