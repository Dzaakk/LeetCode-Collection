package _017
func letterCombinations(digits string) []string {
	result := make([]string, 0)

	if digits == "" {
		return result
	}

	letters := map[byte][]string {
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var backTrack func(index int, combination string)

	backTrack = func(index int, combination string) {
		if len(combination) == len(digits) {
			result = append(result, combination)

			return
		}

		for i := index; i < len(digits); i++ {
			for _, char := range letters[digits[index]] {
				backTrack(i+1, combination+char)
			}
		}
	}
	backTrack(0, "")
	return result
}
