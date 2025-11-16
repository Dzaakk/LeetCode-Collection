package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	seen := map[*ListNode]bool{}

	for head.Next != nil {
		if exists := seen[head]; exists {
			return true
		}

		seen[head] = true
		head = head.Next
	}

	return false
}

// func hasCycle(head *ListNode) bool {
// 	if head == nil || head.Next == nil {
// 		return false
// 	}

// 	slow := head
// 	fast := head.Next

// 	for slow != fast {
// 		if fast == nil || fast.Next == nil {
// 			return false
// 		}
// 		slow = slow.Next
// 		fast = fast.Next.Next
// 	}

// 	return true
// }
