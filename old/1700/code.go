//27:58
func countStudents(students []int, sandwiches []int) int {
	zero, one := 0, 0

	for _, student := range students {
		if student == 0 {
			zero++
		} else {
			one++
		}
	}

	for _, sandwich := range sandwiches {
		if sandwich == 0 {
			if zero == 0 {
				return one
			}
			zero--
		} else {
			if one == 0 {
				return zero
			}
			one--
		}
	}

	return 0
}
