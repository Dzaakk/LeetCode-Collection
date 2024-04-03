//19:35
func reverseWords(s string) string {
	words := strings.Fields(s)
	reversed := ""
	for i := len(words)-1; i >= 0; i--{
		if i == len(words)-1 {
			reversed += words[i]
		} else {
			reversed += " " + words[i]
		}
	}
	return reversed
}

func reverseWords(s string) string {
	var collapsed string
	for i, c := range s {
		if i > 0 && c == ' ' && s[i-1] == ' ' {
			continue
		}
		collapsed += string(c)
	}

	temp := ""
	var listWord []string
	inWord := false
	for _, c := range collapsed {
		if c != ' ' {
			temp += string(c)
			inWord = true
		} else {
			if inWord {
				listWord = append(listWord, temp)
				temp = ""
				inWord = false
			}
		}
	}

	if inWord {
		listWord = append(listWord, temp)
	}

	var words []string
	for _, word := range listWord {
		if word != "" {
			words = append(words, word)
		}
	}
	reversed := ""
	for i := len(words) - 1; i >= 0; i-- {
		if i == len(words)-1 {
			reversed += words[i]
		} else {
			reversed += " " + words[i]
		}
	}

	return reversed
}
