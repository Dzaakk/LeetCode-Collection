package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//18:16
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head
		head = head.Next
		temp.Next = prev
		prev = temp

	}

	return prev
}
