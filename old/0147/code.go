//07:28
func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	for head != nil {
		current := dummy
		for ; current.Next != nil && current.Next.Val < head.Val; current = current.Next {
		}
		current.Next, head.Next, head = head, current.Next, head.Next
	}
	return dummy.Next
}
