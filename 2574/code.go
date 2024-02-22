package _574
//08:00
func leftRightDifference(nums []int) []int {
	var leftSum, rightSum int
	for _, num := range nums {
		rightSum += num
	}

	for i, n := range nums {
		rightSum -= n
		nums[i] = abs(leftSum-rightSum)
		leftSum += n
	}
	return nums

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
