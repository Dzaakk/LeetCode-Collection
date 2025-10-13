func mergeAlternately(word1 string, word2 string) string {
    merged := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		merged += string(word1[i]) + string(word2[j])

		i++
		j++
	}

	merged += word1[i:]
	merged += word2[j:]

	return merged
}