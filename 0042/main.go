func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax, result := 0, 0, 0

	for left <= right {
		if height[left] < height[right] {
			if leftMax > height[left] {
				result = result + (leftMax - height[left])
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if rightMax > height[right] {
				result = result + (rightMax - height[right])
			} else {
				rightMax = height[right]
			}
			right--
		}
	}

	return result
}