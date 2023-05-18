package main

import (
	"fmt"
	"sync"
)
/*
	Gorountine带来的并发问题,用mutex即可解决
*/
func main() {
	var mux sync.Mutex
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				// count++不是原子操作
				// 加上mux让它成“原子操作”
				mux.Lock()
				count++
				mux.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
