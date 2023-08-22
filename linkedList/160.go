package linkedList

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tag := make(map[*ListNode]bool)
	cur := headA
	for cur != nil {
		tag[cur] = true
		cur = cur.Next
	}

	cur = headB
	for cur != nil {
		if tag[cur] {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

func getIntersectionNodev2(headA, headB *ListNode) *ListNode {
	//先分别计算两个的长度，然后计算差值，将两个指针移动到差值为止，然后判断两个节点是否相同
	lenA := 0
	cur := headA
	for cur != nil {
		cur = cur.Next
		lenA++
	}

	lenB := 0
	cur = headB
	for cur != nil {
		cur = cur.Next
		lenB++
	}

	//计算差值
	var lenAB int
	pA, pB := headA, headB
	//移动指针到对齐位置
	if lenA > lenB {
		lenAB = lenA - lenB
		for i := 0; i < lenAB; i++ {
			pA = pA.Next
		}
	} else {
		lenAB = lenB - lenA
		for i := 0; i < lenAB; i++ {
			pB = pB.Next
		}
	}

	//寻找第一个公共节点
	for pA != pB {
		pA = pA.Next
		pB = pB.Next
	}
	return pA
}
