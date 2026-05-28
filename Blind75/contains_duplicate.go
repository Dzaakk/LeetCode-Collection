package main

//01:03
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, n := range nums {
		if _, ok := m[n]; ok {
			return true
		}
		m[n]++
	}

	return false
}
