//12:33
func sortSentence(s string) string {
	words := strings.Fields(s)

	wordIndex := make(map[int]string)

	for _, word := range words {
		index, _ := strconv.Atoi(word[len(word)-1:])
		actualWord := word[:len(word)-1]
		wordIndex[index] = actualWord
	}

	sortedWords := make([]string, len(words))
	for i := 1; i <= len(words); i++ {
		sortedWords[i-1] = wordIndex[i]
	}

	return strings.Join(sortedWords, " ")
}
