func mergeNodes(head *ListNode) *ListNode {
	first := head
	second := head.Next
	for second.Next != nil {
		if second.Val == 0 {
			first.Next = second
			first = first.Next
		} else {
			first.Val += second.Val
		}
		second = second.Next
	}
	first.Next = nil
	return head
}