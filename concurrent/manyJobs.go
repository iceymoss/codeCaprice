package main

import (
	"fmt"
	"net/http"
)

//等待多个 Goroutine 完成：编写一个程序，启动多个 Goroutine 来执行不同的任务，然后等待所有 Goroutine 完成并收集它们的结果。

func jobs1(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}

func jobs2(ch chan int, url string) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	ch <- res.StatusCode
}

func jobs3(ch chan int, value int) {
	ch <- value * value
}

func manyJobs() {
	ch := make(chan int, 3)
	go jobs1([]int{1, 2, 4, 5}, ch)
	go jobs2(ch, "http://81.68.251.185:2000/")
	go jobs3(ch, 10)

	for i := 0; i < 3; i++ {
		data := <-ch
		fmt.Println("返回数据:", data)
	}
	close(ch)
}
