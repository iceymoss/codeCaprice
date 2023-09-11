package stackAndQueue

import "math"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	ans := make([]int, 0)
	l, r := 0, k-1
	for r < len(nums) {
		max := math.MinInt
		for i := l; i <= r; i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		ans = append(ans, max)
		r++
		l++
	}
	return ans
}

type Queue struct {
	queue []int
}

func NewQueue() Queue {
	return Queue{
		make([]int, 0),
	}
}

func (q *Queue) Front() int {
	return q.queue[0]
}

func (q *Queue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *Queue) Empty() bool {
	return len(q.queue) == 0
}

func (q *Queue) Push(val int) {
	for !q.Empty() && val > q.Back() {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, val)
}

func (q *Queue) Pop(val int) {
	if !q.Empty() && val == q.Front() {
		q.queue = q.queue[1:]
	}
}

func maxSlidingWindowV2(nums []int, k int) []int {
	//初始化一个队列
	queue := NewQueue()
	ans := make([]int, 0)

	//将第一个窗口放入队列中
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}

	// 记录前k个元素的最大值
	ans = append(ans, queue.Front())
	for i := k; i < len(nums); i++ {
		//滑动窗口移除最前面的元素
		queue.Pop(nums[k-i])
		//滑动窗口加入最后面的元素
		queue.Push(nums[i])
		ans = append(ans, queue.Front())
	}
	return ans
}
