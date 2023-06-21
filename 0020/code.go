func isValid(s string) bool {

	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
    
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, key := range s {
		if _, ok := pairs[key]; ok {
			stack = append(stack, key)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != key {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}