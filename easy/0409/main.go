func longestPalindrome(s string) int {
	str := make(map[rune]int)
	result := 0
	oddFound := false

	for _, c := range s {
		str[c]++
	}

	for _, count := range str {
		if count%2 == 0 {
			result += count
		} else {
			result += count - 1
			oddFound = true
		}
	}

	if oddFound {
		result++
	}

	return result
}