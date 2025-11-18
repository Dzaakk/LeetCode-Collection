package main

func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[rune]int)
	start, maxLen := 0, 0

	for i, r := range s {
		if prev, ok := lastIndex[r]; ok && prev >= start {
			start = prev + 1
		}

		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}

		lastIndex[r] = i
	}

	return maxLen
}
