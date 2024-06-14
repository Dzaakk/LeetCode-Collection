func deleteDuplicates(head *ListNode) *ListNode {
	tempA := &ListNode{Next: head}
	tempB := tempA

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			tempB.Next = head.Next
		} else {
			tempB = tempB.Next
		}

		head = head.Next
	}

	return tempA.Next
}