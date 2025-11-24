package main

// two pointer
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	i, j := 0, n-1
	idx := n - 1

	for i <= j {
		leftSq := nums[i] * nums[i]
		rightSq := nums[j] * nums[j]

		if leftSq > rightSq {
			res[idx] = leftSq
			i++
		} else {
			res[idx] = rightSq
			j--
		}
		idx--
	}

	return res
}

// sort
// func sortedSquares(nums []int) []int {

// 	for i, n := range nums {
// 		nums[i] = n * n
// 	}

// 	sort.Ints(nums)

// 	return nums
// }
