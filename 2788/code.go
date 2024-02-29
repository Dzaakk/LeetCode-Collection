//22:55
func splitWordsBySeparator(words []string, separator byte) []string {
	res := []string{}
	for _, word := range words {
		var temp string
		for i, w := range word {
			if string(w) == string(separator) {
				if i == 0 || i == len(word)-1 {
					continue
				} else {
					res = append(res, temp)
					temp = ""
					continue
				}
			}
			temp += string(w)
		}
		if temp != ""{
			res = append(res, temp)
		}
	}
	return res
}

//sec
func splitWordsBySeparator(words []string, separator byte) []string {
	var result []string
	for _, word := range words{
		var currentWord []byte
		for i := 0; i < len(word); i++ {
			if word[i] != separator {
				currentWord = append(currentWord, word[i])
			} else if len(currentWord) > 0 {
				result = append(result, string(currentWord))
				currentWord = nil
			}
		}
		if len(currentWord) > 0 {
			result = append(result, string(currentWord))
		}
	}
	return result
}

