//17:18
func detectCapitalUse(word string) bool {
	var counterCaps, counterLows int

	for _, w:= range word {
		if 'a' <= w && w <= 'z' {
			counterLows++
		} else if 'A' <= w && w <= 'Z' {
			counterCaps++
		}
	}

	if len(word) == counterCaps {
		return true
	}
	if counterCaps == 0 {
		return true
	}
	if counterCaps == 1 && 'A' <= word[0] && word[0] <= 'Z' {
		return true
	}

	return false
}
