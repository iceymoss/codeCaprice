package main

import (
	"fmt"
	"sync"
)

//生产者-消费者模型：实现一个生产者-消费者模型，有一个生产者 Goroutine 生成数据并将其发送到通道，多个消费者 Goroutine 从通道接收并处理数据。

func product1(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 100; j++ {
		fmt.Printf("生产者%d生产数据：%d\n", i, j)
		ch <- j
	}
}

func consume1(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("消费者%d消费数据：%d\n", i, v)
		}
	}
}

func productConsume1() {

	wg := &sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go product1(i, ch, wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go consume1(i, ch, wg)
	}

	wg.Wait()
	close(ch)
}
