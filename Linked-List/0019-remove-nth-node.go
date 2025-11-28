package main

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	dummy := &ListNode{0, head}

// 	slow := dummy
// 	fast := dummy
// 	for i := 0; i <= n; i++ {
// 		fast = fast.Next
// 	}

// 	for fast != nil {
// 		fast = fast.Next
// 		slow = slow.Next
// 	}

// 	slow.Next = slow.Next.Next

// 	return dummy.Next
// }

// two pointer
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
