package linkedList

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	//双指针法，n1, n2
	//构造虚拟头节点
	tmpHead := &ListNode{}
	tmpHead.Next = head
	tmp := tmpHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		n1 := tmp.Next
		n2 := tmp.Next.Next
		tmp.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		tmp = n1 // cur移动两位，准备下一轮交换
	}
	return tmpHead.Next
}
