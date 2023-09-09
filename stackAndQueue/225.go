package stackAndQueue

type MyStack struct {
	queue1 []int
	queue2 []int
}

func ConstructorStack() MyStack {
	q1 := make([]int, 0)
	q2 := make([]int, 0)
	return MyStack{
		q1,
		q2,
	}
}

func (this *MyStack) Push(x int) {
	if len(this.queue1) > 0 {
		this.queue1 = append(this.queue1, x)
		return
	}
	if len(this.queue2) > 0 {
		this.queue2 = append(this.queue2, x)
		return
	}
	this.queue1 = append(this.queue1, x)
	return
}

func (this *MyStack) Pop() int {
	//从一个队列中的元素放置到另一个队列中
	if len(this.queue1) == 0 {
		for len(this.queue2) > 0 {
			if len(this.queue2) == 1 { //如果队列只有一个元素
				top := this.queue2[0]
				this.queue2 = this.queue2[1:]
				return top
			}
			//从一个队列中的元素放置到另一个队列中
			top := this.queue2[0]
			this.queue2 = this.queue2[1:]
			this.queue1 = append(this.queue1, top)
		}
	}
	if len(this.queue2) == 0 {
		for len(this.queue1) > 0 {
			if len(this.queue1) == 1 { //如果队列只有一个元素
				top := this.queue1[0]
				this.queue1 = this.queue1[1:]
				return top
			}
			//从一个队列中的元素放置到另一个队列中
			top := this.queue1[0]
			this.queue1 = this.queue1[1:]
			this.queue2 = append(this.queue2, top)
		}
	}
	return 0
}

func (this *MyStack) Top() int {
	var top int
	if len(this.queue1) != 0 {
		top = this.queue1[len(this.queue1)-1]
	}
	if len(this.queue2) != 0 {
		top = this.queue2[len(this.queue2)-1]
	}
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0 && len(this.queue2) == 0
}

func queuePop(q1, q2 *[]int) int {
	for len(*q1) > 0 {
		if len(*q1) == 0 {
			top := (*q1)[0]
			*q1 = (*q1)[1:]
			return top
		}
		top := (*q2)[0]
		*q1 = (*q2)[1:]
		*q2 = append(*q2, top)
	}
	return 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
