//16:24
func wateringPlants(plants []int, capacity int) int {
	steps, cap  := 0, capacity

	for i := 0; i < len(plants); i++ {
		if plants[i] <= cap {
			cap -= plants[i]
			steps++
		} else {
			cap = capacity - plants[i]
			steps += (2 * i) + 1
		}
	}
	return steps
}

