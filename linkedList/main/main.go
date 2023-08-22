package main

import "codeCaprice/linkedList"

func main() {
	link := &linkedList.ListNode{7, &linkedList.ListNode{7, &linkedList.ListNode{7, &linkedList.ListNode{7, nil}}}}
	linkedList.RemoveElements(link, 7)
}
