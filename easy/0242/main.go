func isAnagram(s string, t string) bool {
	runes := make(map[rune]int)

	for _, v := range s {
		runes[v]++
	}

	for _, v := range t {
		runes[v]--
	}

	for _, v := range runes {
		if v != 0 {
			return false
		}
	}

	return true
}