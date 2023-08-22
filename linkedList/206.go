package linkedList

// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	var per *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = per
		per = cur
		cur = next
	}
	return per
}
