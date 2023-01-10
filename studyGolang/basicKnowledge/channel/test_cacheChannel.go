package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 声明一个大小为3的有缓冲的channel

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("sub gorountine结束")

		for i := 0;i < 4;i++ {
			c <- i // 把i放到channel中
			fmt.Println("sub gorountine正在运行，发送的元素为: ", i, "len(c) = ", len(c), "cap(c) = ", cap(c))
		}
	}()
	
	time.Sleep(1 * time.Second)
	for i := 0;i < 4;i++ {
		num := <-c //从channel c中取数据并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main gorountine结束")
}