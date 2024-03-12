//05:33
func countConsistentStrings(allowed string, words []string) int {
	letters := make([]bool, 26)
	res := 0

	for _, ch := range allowed {
		letters[ch - 'a'] = true
	}

	for i := 0; i < len(words); i++ {
		for j, ch := range words[i] {
			if !letters[ch - 'a'] {
				break
			}
			if j == len(words[i]) - 1{
				res++
			}
		}
	}
	return res
}

