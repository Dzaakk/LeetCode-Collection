//13:15
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	sort.Ints(nums)
	dummy := &ListNode{}
	newHead := dummy
	for _, num:= range nums{
		newHead.Next = &ListNode{Val: num}
		newHead = newHead.Next
	}

	return dummy.Next
}