package _367
//11:47
func arithmeticTriplets(nums []int, diff int) int {
	counter := 0
	numCount := make(map[int]int)

	for _, num := range nums {
		numCount[num]++
	}

	for num := range numCount {
		if numCount[num+diff] > 0 && numCount[num+2*diff]>0 {
			counter++
		}
	}
	return counter
}