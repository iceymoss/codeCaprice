package main

import (
	"fmt"
	"sync"
)

//Goroutine 并发计算：编写一个程序，启动多个 Goroutine 来计算一系列整数的平方，并将结果存储在共享的数据结构中。最后，将结果打印出来。

func square(num int, ch chan int) {
	ch <- num * num
}

func chanSquare() {
	nums := []int{1, 2, 4, 5, 6, 7, 7}
	ch := make(chan int)
	defer close(ch)
	for _, v := range nums {
		go square(v, ch)
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch)
	}
}

func mutexSquare(nums []int) {
	ans := make([]int, 0)
	m := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(v int, m *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			m.Lock()
			ans = append(ans, nums[v]*nums[v])
			m.Unlock()
		}(i, m, wg)
	}
	wg.Wait()
	fmt.Println(ans)
}
