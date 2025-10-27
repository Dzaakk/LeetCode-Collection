package main

func containsDuplicate(nums []int) bool {
	mapNum := make(map[int]int)

	for _, n := range nums {
		if _, ok := mapNum[n]; ok {
			return true
		}
		mapNum[n]++
	}
	return false
}
