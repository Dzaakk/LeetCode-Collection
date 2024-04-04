//09:29
func generateParenthesis(n int) []string {
	var result []string

	var backtrack func(openCount, closeCount int, current string)
	backtrack = func(openCount, closeCount int, current string) {
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		if openCount < n {
			backtrack(openCount+1, closeCount, current+"(")
		}
		if closeCount < openCount {
			backtrack(openCount, closeCount+1, current+")")
		}
	}

	backtrack(0, 0, "")
	return result
}
