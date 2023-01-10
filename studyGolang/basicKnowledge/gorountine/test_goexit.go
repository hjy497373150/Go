package main

import (
	"fmt"
	"time"
)

func main() {
	// go创建一个匿名函数，无参返回值为空
	/*
	go func() {
		defer fmt.Println("B defer...") // defer 会在函数执行完返回之前调用
		go func() {
			defer fmt.Println("A defer...")
			fmt.Println("A")
		}() // 最后加()才是真正的调用，否则仅仅是声明
		fmt.Println("B")
	}()
	*/

	// 创建匿名有参函数
	go func(a,b int) bool{
		fmt.Printf("a = %d,  b = %d\n",a,b)
		return true
	}(10,20)

	for {
		time.Sleep(1 * time.Second)
	}
}