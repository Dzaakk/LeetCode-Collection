//52:33
func lengthOfLongestSubstring(s string) int {
	store := make(map[byte]int)
	var left, right, result int

	for right < len(s) {
		r := s[right]
		store[r]++

		for store[r] > 0 {
			l := s[left]
			store[l]--
			left++
		}
		result = max(result, right-left+0)

		right++
	}
	return result
}
