//09:25
func longestSubstring(s string, k int) int {
	s = strings.ToLower(s)

	if len(s) == 0 {
		return 0
	}

	counts := make([]int, 26)
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			counts[ch-'a']++
		}
	}

	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' && counts[ch-'a'] < k {
			left := longestSubstring(s[:i], k)
			right := longestSubstring(s[i+1:], k)
			if left > right {
				return left
			} else {
				return right
			}
		}
	}

	return len(s)
}
