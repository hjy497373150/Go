package main

import (
	"fmt"
	"time"
)

// 从gorountine
func newRountine() {
	i := 0
	for {
		i++
		fmt.Printf("newRountine's i = %d\n",i)
		time.Sleep(1*time.Second)
	}
}

// 主gorountine
func main() {
	// 创建一个新的 goroutine
	go newRountine()

	fmt.Println("exit main goRountine...") // 主gorountine退出，则从goroutine也结束
	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("mainRountine's i = %d\n",i)
	// 	time.Sleep(1*time.Second)
	// }
}