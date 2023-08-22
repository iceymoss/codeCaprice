package linkedList

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	//使用哈希表来记录
	tagMpa := make(map[*ListNode]int)
	cur := head
	index := 0
	for cur != nil {
		if _, ok := tagMpa[cur]; !ok {
			tagMpa[cur] = index
		} else {
			return cur
		}
		index++
		cur = cur.Next
	}
	return nil
}

func detectCycleV2(head *ListNode) *ListNode {
	//双指针，快慢指针
	var fast, slow *ListNode
	fast = head
	slow = head
	for fast != nil && fast.Next != nil {
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return head
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
