package main

//02:18
func hasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]int)

	if head == nil {
		return false
	}

	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}

		seen[head]++

		head = head.Next
	}

	return false
}
