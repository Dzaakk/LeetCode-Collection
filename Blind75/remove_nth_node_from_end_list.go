package main

//05:27
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return head
}
