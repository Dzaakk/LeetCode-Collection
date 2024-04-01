//03:52
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var counter int
	for _, cost := range costs {
		if cost > coins {
			break
		} else {
			coins -= cost
			counter++
		}
	}
	return counter
}
