package linkedList

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	dummyHead *Node //虚拟头节点
	Size      int   //链表大小
}

func Constructor() MyLinkedList {
	node := &Node{-999, nil}

	return MyLinkedList{
		node,
		0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	//判断索引是否有效
	if this == nil || this.Size <= index || index < 0 {
		return -1
	}
	// 让cur等于真正头节点
	cur := this.dummyHead.Next
	for i := 0; i < index; i++ {
		//i < index遍历到最后cur = cur.Next =》正好 i == index
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	//构造节点
	node := &Node{}
	node.Val = val
	node.Next = this.dummyHead.Next
	this.dummyHead.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	//追加到尾部,先找到当前最后一个节点
	cur := this.dummyHead
	for cur.Next != nil {
		cur = cur.Next
	}
	node := &Node{Val: val}
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	//判断索引是否有效
	if index < 0 {
		index = 0
	} else if index > this.Size {
		return
	}

	//添加节点
	node := &Node{Val: val}
	cur := this.dummyHead
	//遍历到需要添加的节点处
	for i := 0; i < index; i++ {
		cur = cur.Next //这里以虚拟节点为0,然后i < index和cur = cur.Next正好可以找到index前一个节点
	}
	node.Next = cur.Next
	cur.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//判断index是否有效, 链表索引从0开始的，所以index = this.Size也应该失效的
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next // 遍历到要删除节点的前一个节点
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	this.Size--
}
