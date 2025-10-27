package main

func twoSum(nums []int, target int) []int {
	mapNum := make(map[int]int)

	for i, n := range nums {
		if j, ok := mapNum[n]; ok {
			return []int{i, j}
		}
		dif := target - n
		mapNum[dif] = i

	}
	return []int{}
}
