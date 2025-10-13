func findCenter(edges [][]int) int {
	var temp int
	var count = 1
	for i := 0; i < len(edges[i]); i++ {
		temp = edges[0][i]
		for j := 1; j < len(edges); j++ {

			for k := 0; k < len(edges[j]); k++ {
				if temp == edges[j][k] {
					count++
				}
				if count == len(edges) {
					return temp
				}
			}
		}
	}
	return temp
}