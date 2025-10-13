func getDecimalValue(head *ListNode) int {
	decimal := 0

	for head != nil {
		decimal = decimal*2 + head.Val
		head = head.Next
	}

	return decimal
}
