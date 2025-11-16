package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}

//recursive
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}

// 	if list2 == nil {
// 		return list1
// 	}

// 	if list1.Val <= list2.Val {
// 		list1.Next = mergeTwoLists(list1.Next, list2)
// 		return list1
// 	}
// 	list2.Next = mergeTwoLists(list2.Next, list1)
// 	return list2
// }
