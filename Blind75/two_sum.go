package main

//06:31
func twoSum(nums []int, target int) []int {
	mapNum := make(map[int]int)

	for i, n := range nums {
		x := target - n

		if j, ok := mapNum[n]; ok {
			return []int{j, i}
		}

		mapNum[x] = i

	}
	return []int{0, 0}
}
