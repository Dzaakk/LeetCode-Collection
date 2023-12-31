/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next = deleteDuplicates(head.Next); head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}