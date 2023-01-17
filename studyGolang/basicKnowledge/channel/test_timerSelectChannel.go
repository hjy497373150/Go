package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()

	// 每次运行结果可能都不同，select每次只选择一个来执行
	select {
	case num1:= <-ch1:
		fmt.Println("ch1中取数据:", num1)
	case num2,ok := <-ch2:
		if ok {
			fmt.Println("ch2中取数据:",num2)
		} else {
			fmt.Println("ch2已关闭！")
		}
	// 超时退出，可设置
	case <-time.After(time.Second * 5):
        fmt.Println("request time out")
	}

}