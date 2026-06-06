package main

//38:47
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)

	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		m[s[r]]++

		for m[s[r]] > 1 {
			m[s[l]]--
			l++
		}

		currLen := r - l + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}
