//08:54
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	mapLeft := make(map[string]uint8)
	mapRight := make(map[uint8]string)

	for i, word := range words {
		if val, ok := mapLeft[word]; ok && val != pattern[i] {
			return false
		}

		if val, ok := mapRight[pattern[i]]; ok && val != word {
			return false
		}

		mapLeft[word] = pattern[i]
		mapRight[pattern[i]] = word
	}

	return true
}
