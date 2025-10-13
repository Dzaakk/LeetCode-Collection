//05:41
func repeatedNTimes(nums []int) int {
	m := make(map[int]int)
	var duplicate int
	for _, num := range nums {
		m[num]++
		if i := m[num]; i > 1 {
			duplicate = num
			break
		}
	}
	return duplicate
}