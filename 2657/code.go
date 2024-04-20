//06:26
func findThePrefixCommonArray(A []int, B []int) []int {
	mapA := make(map[int]bool, 0)
	mapB := make(map[int]bool, 0)
	counter := 0

	answer := make([]int, 0)

	for i, n := range A {
		mapA[n] = true
		mapB[B[i]] = true

		if mapB[n] {
			counter++
		}
		if mapA[B[i]] {
			counter++
		}
		if n == B[i] {
			counter--
		}

		answer = append(answer, counter)
	}

	return answer
}