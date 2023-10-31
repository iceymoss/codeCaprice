package main

import (
	"fmt"
	"sync"
)

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
