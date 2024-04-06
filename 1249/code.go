//10:02
func minRemoveToMakeValid(s string) string {
	var stack []int
	chars := []rune(s)

	for i, char := range chars {
		if char == '(' {
			stack = append(stack, i)
		} else if char == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				chars[i] = 0
			}
		}
	}

	for _, idx := range stack {
		chars[idx] = 0
	}

	result := make([]rune, 0 , len(chars))
	for _, char := range chars {
		if char != 0 {
			result = append(result, char)
		}
	}

	return string(result)
}
