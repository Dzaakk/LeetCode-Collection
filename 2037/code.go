//03:53
func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var moves int
	for i := 0; i < len(seats); i++ {
		if seats[i] > students[i] {
			moves += seats[i] - students[i]
		} else {
			moves += students[i] - seats[i]
		}
	}

	return moves
}