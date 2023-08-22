package linkedList

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//构造虚拟节点
	tmp := &ListNode{}
	tmp.Next = head
	cur := tmp.Next

	//统计总长度
	count := 0
	for cur != nil {
		cur = cur.Next
		count++
	}

	//例如：n = 2; count = 4, 4-2 = 2， 有虚拟节点开始：0-1->2正好在要移除的节点的上一个节点
	count -= n
	cur = tmp
	for i := 0; i < count; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	return tmp.Next
}
