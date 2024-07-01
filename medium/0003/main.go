func lengthOfLongestSubstring(s string) int {
	runeMap := map[rune]int{}
	maxLen := 0

	j := -1
	for i, v := range s {
		index, ok := runeMap[v]
		if ok {
			j = max(j, index)
		} else {
			j = max(j, -1)
		}

		maxLen = max(maxLen, i-j)
		runeMap[v] = i
	}

	return maxLen
}
