package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func write(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num := rand.Intn(101)
		fmt.Printf("生产数据%d生产了数据：%d\n", i, num)
		ch <- num
		time.Sleep(1 * time.Second)
	}
}

func read(i int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("消费者%d消费了数据：%d\n", i, v)
	}
}

// 请使用Go语言实现一个生产者消费者模型。你需要创建多个生产者和消费者，它们并发地运行。生产者生产数据并将其发送到channel中，消费者从channel中读取数据并处理。
// 具体要求：
// 1. 创建3个生产者goroutine， 每个生产者每隔1秒产生一个0-100的随机整数，然后将其发送到channel中。
// 2.创建2个消费者goroutine， 它们不断地人channel中读取数据并打印。
// 3. 为了避免channel无限增长，你需要设置一个合适的buffer size。
// 4. 主程序应当在所有生产者和消费者都完成工作后结束。

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	for j := 0; j < 2; j++ {
		wg.Add(1)
		go read(j, ch, &wg)
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write(i, ch, &wg)
	}
	wg.Wait()
}
