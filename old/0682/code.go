func calPoints(operations []string) int {
	stack := []int{}

	for _, op := range operations {
		switch op {
		case "+":
			if len(stack) >= 2 {
				newScore := stack[len(stack)-1] + stack[len(stack)-2]
				stack = append(stack, newScore)
			}
		case "D":
			if len(stack) >= 1 {
				newScore := 2 * stack[len(stack)-1]
				stack = append(stack, newScore)
			}
		case "C":
			if len(stack) >= 1 {
				stack = stack[:len(stack)-1]
			}
		default:
			score, _ := strconv.Atoi(op)
			stack = append(stack, score)
		}
	}

	total := 0
	for _, score := range stack {
		total += score
	}

	return total
}