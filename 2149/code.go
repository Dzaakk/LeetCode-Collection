//05:03
func rearrangeArray(nums []int) []int {
	var right, left []int

	for _, num := range nums {
		if num > 0 {
			right = append(right, num)
		} else {
			left = append(left, num)
		}
	}

	rightIdx := 0
	leftIdx:= 0
	for i := 0; i < len(nums); i++ {
		if i % 2 == 0 {
			nums[i] = right[rightIdx]
			rightIdx++
		} else {
			nums[i] = left[leftIdx]
			leftIdx++
		}
	}

	return nums
}
