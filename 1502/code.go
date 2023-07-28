func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	dif := arr[1] - arr[0]

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != dif {
			return false
		}
	}

	return true
}