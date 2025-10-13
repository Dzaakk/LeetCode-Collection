//08:09
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		counter1 := bits.OnesCount(uint(arr[i]))
		counter2 := bits.OnesCount(uint(arr[j]))
		return counter1 < counter2 || (counter1 == counter2 && arr[i] < arr[j])
	})
	return arr
}

