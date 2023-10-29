package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 时间复杂度：O(n), 空间复杂度：O(n)
func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}

	//双指针
	l, r := 0, len(stack)-1
	for l <= r {
		if stack[l] != stack[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// 时间复杂度：O(n), 空间复杂度：O(1)
func isPalindromeV1(head *ListNode) bool {
	cur := head
	lentgh := 0
	for cur != nil {
		lentgh += 1
		cur = cur.Next
	}

	cur = head
	tmp := 0
	for cur != nil {
		if tmp == lentgh/2 {
			tow := revers(cur)
			cur = head
			for tow != nil && cur != nil {
				if tow.Val != cur.Val {
					return false
				}
				tow = tow.Next
				cur = cur.Next
			}
			return true
		}
		tmp++
		cur = cur.Next
	}

	return true
}

func revers(head *ListNode) *ListNode {
	//构造头结点
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
