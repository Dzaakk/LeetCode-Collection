//06:36
func checkIfPangram(sentence string) bool {
	store := make(map[rune]bool)
	for _, char := range sentence {
		store[char] = true

	}
	return len(store)==26
}