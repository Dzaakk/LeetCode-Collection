//0825
func firstUniqChar(s string) int {
	counter := make(map[rune]int)
	for _, c := range s{
		counter[c]++
	}

	for i, c := range s {
		if counter[c] == 1 {
			return i
		}
	}
	return -1

}
