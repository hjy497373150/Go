package main

import "fmt"

func fibonachi(c, quit chan int) {
	x,y := 1,1
	// select 可以监控多channel的状态
	for {
		select {
		case c <- x:
			// 如果可以把x写到c中 执行
			t := x
			x = y
			y = t + y
		case <-quit:
			// 如果quit可读
			fmt.Println("sub gorountine exit...")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub gorountine
	go func() {
		for i := 0;i < 10;i++ {
			fmt.Println(<-c) //打印c中的内容
		}
		quit <- 0 // 把0发送到quit channel
	}()
	
	// main gorountine
	fibonachi(c, quit)

}