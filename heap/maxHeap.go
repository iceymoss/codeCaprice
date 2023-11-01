package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{
		make([]int, 1),
	}
}

func (h *MaxHeap) Insert(value int) {
	h.arr = append(h.arr, value)
	index := len(h.arr) - 1

	//h.shiftUp(index)

	for index/2 > 0 && h.arr[index] > h.arr[index/2] {
		h.arr[index], h.arr[index/2] = h.arr[index/2], h.arr[index]
		index /= 2
	}
}

func (h *MaxHeap) Pop() int {
	index := 1
	max := h.arr[index]
	h.arr[index] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	//h.shiftDown(index)
	for 2*index < len(h.arr) { //左孩子是否存在
		j := 2 * index
		if j+1 < len(h.arr) && h.arr[j] < h.arr[j+1] { //右孩子
			j += 1
		}
		if h.arr[index] > h.arr[j] {
			break
		}
		h.arr[index], h.arr[j] = h.arr[j], h.arr[index]
		index = j
	}
	return max
}

func (h *MaxHeap) shiftDown(i int) {
	for 2*i < len(h.arr) {
		j := 2 * i
		if j+1 < len(h.arr) && h.arr[j] < h.arr[j+1] {
			j++
		}
		if h.arr[i] > h.arr[j] {
			break
		}
		h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
		i = j
	}
}

func (h *MaxHeap) shiftUp(i int) {
	for i/2 > 0 && h.arr[i] > h.arr[i/2] {
		h.arr[i], h.arr[i/2] = h.arr[i/2], h.arr[i]
		i /= 2
	}
}

func (h *MaxHeap) heapSort(arr []int) []int {
	ans := make([]int, len(arr))
	for _, v := range arr {
		h.Insert(v)
	}
	for i := 0; i < len(arr); i++ {
		ans[i] = h.Pop()
	}
	return ans
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	maxHeap := NewMaxHeap()
	arr := []int{1, 2, 3, 4, 5, 6, 100, 34, 64, 32, 56, 34, 77, 0, 31}
	brr := append([]int{}, arr...)
	fmt.Println(brr)
	fmt.Println(maxHeap.heapSort(arr))
	fmt.Println(bubbleSort(arr))
}
