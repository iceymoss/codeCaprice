package main

import (
	"fmt"
	"sync"
)

//互斥锁：创建一个共享的整数变量，并启动多个 Goroutine 来递增该变量。使用互斥锁（sync.Mutex）来确保每个 Goroutine 修改变量时不会发生竞争条件。

func shareValue() {
	var value int
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v *int, wg *sync.WaitGroup, m *sync.Mutex) {
			defer wg.Done()
			m.Lock()
			*v++
			m.Unlock()
		}(&value, wg, m)
	}
	wg.Wait()
	fmt.Println(value)
}
