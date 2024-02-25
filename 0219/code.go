//11:01
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		val, ok := m[num]
		if ok{
			if i - val <= k {
				return true
			}
		}
		m[num] = i
	}
	return false
}
