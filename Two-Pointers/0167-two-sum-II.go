package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		num := numbers[l] + numbers[r]

		if num < target {
			l++
		}
		if num > target {
			r--
		}
		if num == target {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
}
