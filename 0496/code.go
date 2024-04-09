//36:59
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	nextGreater := make(map[int]int)
	stack := make([]int, 0)

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			nextGreater[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	for i, num := range nums1 {
		if next, ok := nextGreater[num]; ok {
			result[i] = next
		} else {
			result[i] = -1
		}
	}
	return result
}
