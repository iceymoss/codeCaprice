package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	//输入数组
	var arr []int
	for {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			break
		}
		arr = append(arr, num)
	}

	fmt.Println(arr)

	ans := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(arr)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index <= len(arr) {
			ans = append(ans, append([]int{}, path...))
		}

		if index > len(arr) {
			return
		}

		for i := index; i < len(arr); i++ {
			if i != index && arr[i] == arr[i-1] {
				continue
			}
			path = append(path, arr[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	fmt.Println(ans)

	//输入字符串
	//var str string
	//fmt.Scanf("%s", &str)
	//
	//fmt.Println(str)
	//
	////输入链表
	//input := "1->2->3->4"
	//parts := strings.Split(input, "->")
	//
	//var head, prev *ListNode
	//
	//for _, valStr := range parts {
	//	val, _ := strconv.Atoi(valStr)
	//	node := &ListNode{Val: val}
	//	if prev == nil {
	//		head = node
	//	} else {
	//		prev.Next = node
	//	}
	//	prev = node
	//}
	//
	//printList(head)

}
