func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	counter := 0
	for _, h := range hours {
		if h >= target {
			counter++
		}
	}

	return counter
}