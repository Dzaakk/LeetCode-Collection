func magicalString(n int) int {
	if n <= 3 {
		return 1
	}

	magic, idx := make([]int, n), 3
	occur, value, count := 2, 1, 1
	magic[0], magic[1], magic[2] = 1, 2, 2

	for idx < n {
		for i := 0; i < magic[occur] && idx < n; i++ {
			magic[idx] = value
			if value == 1 {
				count++
			}
			idx++
		}
		value ^= 3
		occur++
	}

	return count
}