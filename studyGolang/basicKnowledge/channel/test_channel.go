package main 

import (
	"fmt"
)
func main() {
	// 创建一个channel 无缓冲
	c := make(chan int) 
	go func() {
		defer fmt.Printf("gorountine结束...\n")
		fmt.Println("gorountine正在运行...")

		c <- 666 // 将666发送给channel c
	} ()
	num := <- c  // 把channel c 中的内容赋给num
	// 保证了调用的先后顺序，因为666必须要写到c中才能赋给num 即num := <- c必在c <- 666之后
	fmt.Println("num = ",num)
	fmt.Println("mainrountine结束...")
}