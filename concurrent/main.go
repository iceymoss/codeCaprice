package main

/*
Goroutine 并发计算：编写一个程序，启动多个 Goroutine 来计算一系列整数的平方，并将结果存储在共享的数据结构中。最后，将结果打印出来。

互斥锁：创建一个共享的整数变量，并启动多个 Goroutine 来递增该变量。使用互斥锁（sync.Mutex）来确保每个 Goroutine 修改变量时不会发生竞争条件。

生产者-消费者模型：实现一个生产者-消费者模型，有一个生产者 Goroutine 生成数据并将其发送到通道，多个消费者 Goroutine 从通道接收并处理数据。

等待多个 Goroutine 完成：编写一个程序，启动多个 Goroutine 来执行不同的任务，然后等待所有 Goroutine 完成并收集它们的结果。

定时任务：编写一个程序，定期执行某个任务，例如每隔一秒钟打印一次当前时间。可以使用 time.Tick 来触发定时任务。

信号量模式：使用信号量来控制一组 Goroutine 的并发访问。例如，实现一个连接池，限制同时可以使用的连接数。

并发安全数据结构：创建一个并发安全的数据结构，如并发安全的队列、栈或映射，以便多个 Goroutine 可以安全地读写数据。

并发文件读写：编写一个程序，启动多个 Goroutine 来同时读取和写入文件。确保读写操作不会产生竞争条件。
*/

func main() {
	//chanSquare()
	//mutexSquare([]int{1, 2, 4, 5, 6, 7, 7})
	//
	//shareValue()

	//productConsume0()

	//productConsume1()

	//productConsume3()

	//manyJobs()
	productConsume3()
}
