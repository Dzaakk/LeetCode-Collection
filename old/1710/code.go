//10:55
func maximumUnits(boxTypes [][]int, truckSize int) int {
	var result int

	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	for i := 0; i < len(boxTypes); i++ {
		for boxTypes[i][0] > 0 && truckSize > 0 {
			result += boxTypes[i][1]
			boxTypes[i][0]--
			truckSize--
		}
	}
	return result
}