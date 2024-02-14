//27:52
func canCompleteCircuit(gas []int, cost []int) int {
	start, tank, total := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank = tank - cost[i] + gas[i]
		total = total - cost[i] + gas[i]

		if tank < 0 {
			start = i +1
			tank = 0
		}
	}

	if total >= 0 {
		return start
	}

	return -1
}