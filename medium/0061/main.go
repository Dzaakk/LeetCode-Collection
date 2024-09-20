func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 1
	curr := head

	for curr.Next != nil {
		length++
		curr = curr.Next
	}

	curr.Next = head

	k %= length
	lengthPrev := length - k - 1
	curr = head

	for i := 0; i < lengthPrev; i++ {
		curr = curr.Next
	}

	res := curr.Next

	curr.Next = nil

	return res
}