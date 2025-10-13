package _349
//05:38
func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)
	set2:= make(map[int]bool)

	for _, num := range nums1 {
		set1[num] = true
	}

	for _, num := range nums2 {
		if set1[num] {
			set2[num] = true
		}
	}

	result := make([]int, 0, len(set2))
	for num, _ := range set2 {
		result = append(result, num)
	}
	return result
}