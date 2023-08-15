// loop
func removeElements(head *ListNode, val int) *ListNode {
	var result *ListNode

	for head != nil {
		if head.Val == val {
			head = head.Next
		} else {
			result = head
			break
		}
	}

	current := result
	for current != nil && current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return result
}

// recursive
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	}
	head.Next = removeElements(head.Next, val)
	return head
}