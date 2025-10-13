func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		a, b := current.Next.Val, current.Next.Next.Val
		gcdValue := GCD(a, b)

		newNode := &ListNode{Val: gcdValue}
		newNode.Next = current.Next.Next
		current.Next.Next = newNode

		current = current.Next.Next
	}

	return dummy.Next
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

