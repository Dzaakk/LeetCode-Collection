//14:21
func minimumBoxes(apple []int, capacity []int) int {
	sort.Ints(capacity)

	sum := 0
	for _, a := range apple {
		sum += a
	}

	n := len(capacity)
	for i := n - 1; i >= 0; i--{
		if sum <= capacity[i] {
			return n - i
		}

		sum -= capacity[i]
	}
	return n
}