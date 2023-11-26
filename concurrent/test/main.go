package main

import (
	"fmt"
	"sync"
)

//生产者者-消费者模型：实现3个生产者-10个消费者模型，有3个生产者 Goroutine 生成数据并将其发送到通道，10个消费者 Goroutine 从通道接收并处理数据。然后结束程序

func producer3(id int, ch chan int, wg *sync.WaitGroup, msgs chan struct{}) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		date := id*10 + i
		ch <- date
		fmt.Printf("生产者%d生产数据：%d\n", id, date)
	}
	//当前生生产者结束，通知消费者
	msgs <- struct{}{}
}

func consumer3(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("消费者%d消费数据：%d\n", id, v)
	}
}

func getMsgToCloseCh(chs chan struct{}, ch chan int, count int) {
	//处理生产者状态
	for i := 0; i < count; i++ {
		<-chs //如果生产者没有发送状态，说明在生产，chs缓存为producerCount，如果有就读，没有当前协程堵塞
	}

	//当所有生产者结束后，关闭主ch，从而唤醒所以消费者
	close(ch)
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	producerCount := 3

	//传输每一个生产者的状态
	chMsg := make(chan struct{}, producerCount)

	for i := 0; i < producerCount; i++ {
		wg.Add(1)
		go producer3(i, ch, wg, chMsg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go consumer3(i, ch, wg)
	}

	go getMsgToCloseCh(chMsg, ch, producerCount)

	wg.Wait()
}

//func rervire(array []int, begin, end int) []int {
//	if begin >= end || begin < 0 || end >= len(array) {
//		return array
//	}
//	for begin < end {
//		array[begin], array[end] = array[end], array[begin]
//		begin++
//		end--
//	}
//	return array
//}
