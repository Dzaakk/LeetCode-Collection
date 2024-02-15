//11.54
func majorityElement(nums []int) int {
	counts := map[int]int{}

	for _, num := range nums {
		counts[num]++
	}
	majority := 0
	maxCount:= 0
	for num, count := range counts{
		if count > maxCount{
			majority = num
			maxCount = count
		}
	}
	return majority
}

func majorityElement(nums []int) int {
	var ans int
	var count int

	for _, num := range nums {
		if count == 0 {
			ans = num
		}
		if num == ans {
			count++
		} else {
			count--
		}
	}

	return ans
}