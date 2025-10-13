//42:42
func maxFrequencyElements(nums []int) int {
	mapCount := make(map[int]int)
	max, temp :=  0, 0
	for _, n := range nums {
		mapCount[n]++
		if mapCount[n] > temp {
			temp = mapCount[n]
		}
	}
	for _, v := range mapCount {
		if v == temp {
			max += temp
		}
	}
	return max
}