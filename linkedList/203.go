package linkedList

func RemoveElements(head *ListNode, val int) *ListNode {
	cur := head
	per := &ListNode{}
	per.Next = head
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return per.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	ans := cur
	per := &ListNode{}
	per.Next = cur
	for cur != nil {
		if cur.Val == val {
			per.Next = cur.Next
			cur = cur.Next
		} else {
			per = per.Next
			cur = cur.Next
		}
	}
	return ans
}
