package main

//11:09
func missingNumber(nums []int) int {
	mapNum := make(map[int]int)
	for _, n := range nums {
		mapNum[n]++
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := mapNum[i]; !ok {
			return i
		}
	}
	return 0
}

// better solution
func missingNumber2(nums []int) int {
	l := len(nums)

	full := l * (l + 1) / 2
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	return full - actualSum
}
