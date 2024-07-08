func canConstruct(ransomNote string, magazine string) bool {

	counter := len(ransomNote)
	letters := make(map[rune]int)

	for _, c := range ransomNote {
		letters[c]++
	}

	for _, c := range magazine {
		if ok := letters[c]; ok != 0 {
			letters[c]--
			counter--
			if counter == 0 {
				return true
			}

		}
	}
	return false
}