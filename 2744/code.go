//24:42
func maximumNumberOfStringPairs(words []string) int {
	counter := 0
	for i := 0; i < len(words)-1; i++ {
		rev := ""
		word := string(words[i])
		for i := len(word) -1; i >= 0; i--{
			rev += string(word[i])
		}
		for j := i+1; j < len(words); j++ {
			if rev == string(words[j]){
				counter++
				break
			}
		}
	}
	return counter
}
