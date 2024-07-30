func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next

		first.Next = second.Next
		second.Next = first
		current.Next = second

		current = first
	}
	return dummy.Next
}