//19:00
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	result := make([]int, 2)
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	for _, n := range nums1{
		m1[n] = true
	}
	for _, m  := range nums2{
		m2[m] = true
	}

	for _, num := range nums1 {
		if m2[num] {
			result[0]++
		}
	}
	for _, num := range nums2 {
		if m1[num] {
			result[1]++
		}
	}
	return result
}