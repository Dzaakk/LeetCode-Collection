//20:42
func timeRequiredToBuy(tickets []int, k int) int {
	time := 0
	totalTickets := 0
	for i := 0; i < len(tickets) ; i++ {
		totalTickets += tickets[i]
	}
	for tickets[k] > 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] > 0 {
				tickets[i]--
				time++
				if i == k && tickets[k] == 0 {
					return time
				}
			}
		}
	}
	return time
}

