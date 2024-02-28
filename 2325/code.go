package _325
//19:40
func decodeMessage(key string, message string) string {
	mapKey := make(map[string]int)

	keys := []string{}
	alfabet:= []string{}

	i:= 0
	for _, s := range key{
		if s != ' '{
			cek := string(s)
			mapKey[cek]++
			if mapKey[cek] == 1 {
				sum := 'a' + i
				keys = append(keys, cek)
				alfabet = append(alfabet, string(sum))
				i++
			}
		}
	}

	result := ""
	for i := 0 ; i <len(message); i++ {
		for j:= 0; j< len(keys); j ++ {
			if string(message[i]) == keys[j]{
				result += alfabet[j]
			} else if message[i] == ' '{
				result += " "
				break
			}
		}
	}
	return result
}


//sec way
func decodeMessage(key string, message string) string {
	maps := make(map[rune]rune)
	alfabet := 'a'

	for _, character := range key {
		if character == ' '{
			continue
		}

		if _, ok := maps[character]; !ok {
			maps[character] = alfabet
			alfabet = rune(alfabet+1)
		}
	}

	result := ""

	for _, char := range message {
		if char == ' ' {
			result += " "
			continue
		}
		result += string(maps[char])
	}
	return result
}