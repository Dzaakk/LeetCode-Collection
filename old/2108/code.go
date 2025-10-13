//05:56
func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}
func isPalindrome(word string) bool {
	n := len(word)
	for i := 0; i < n/2; i++ {
		if word[i] != word[n-i-1]{
			return false
		}
	}
	return true
}
//2
func firstPalindrome(words []string) string {
	for _, word := range words {
		check := true
		for i := 0; i < (len(word)/2); i++ {
			if word[i] != word[len(word)-i-1] {
				check = false
				break
			}
		}

		if check == true {
			return word
		}
	}

	return ""
}