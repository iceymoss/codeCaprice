package stackAndQueue

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	return MyQueue{
		s1,
		s2,
	}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	//s1 -> s2
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	top := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return top
}

func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		return this.stack1[0]
	} else {
		return this.stack2[len(this.stack2)-1]
	}
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}
