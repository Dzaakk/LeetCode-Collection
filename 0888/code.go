//18:06
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA, sumB := 0, 0
	aliceNums, bobNumns := make(map[int]struct{}), make(map[int]struct{})
	for _, a := range aliceSizes {
		sumA += a
		aliceNums[a] = struct{}{}
	}
	for _, b := range bobSizes{
		sumB += b
		bobNumns[b] = struct{}{}
	}

	for n, _ := range aliceNums {
		x := ((sumB - sumA) / 2) + n
		if _, found := bobNumns[x]; found {
			return []int{n, x}
		}
	}
	return []int{}
}
