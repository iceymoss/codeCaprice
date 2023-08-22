package main

import (
	"codeCaprice/linkedList"
	"fmt"
)

func main() {
	link := &linkedList.ListNode{7, &linkedList.ListNode{7, &linkedList.ListNode{7, &linkedList.ListNode{7, nil}}}}
	linkedList.RemoveElements(link, 7)
	fmt.Println("git")
	linkedList.Constructor()
}
