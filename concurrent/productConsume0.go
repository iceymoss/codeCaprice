package main

import (
	"fmt"
	"sync"
)

//生产者-消费者模型：实现一个生产者-消费者模型，有一个生产者 Goroutine 生成数据并将其发送到通道，多个消费者 Goroutine 从通道接收并处理数据。

func product0(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	tag := 0
	for {
		if tag == 100 {
			close(ch)
			return
		}
		fmt.Printf("生产者生产数据：%d\n", tag)
		ch <- tag
		tag++
	}
}

func consume0(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("消费者%d消费数据：%d\n", i, v)
	}
}

func productConsume0() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go product0(ch, wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go consume0(i, ch, wg)
	}
	wg.Wait()
}
