func partition(head *ListNode, x int) *ListNode {
	leftDummy := &ListNode{}
	rightDummy := &ListNode{}

	left := leftDummy
	right := rightDummy

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}

	right.Next = nil
	left.Next = rightDummy.Next

	return leftDummy.Next

}