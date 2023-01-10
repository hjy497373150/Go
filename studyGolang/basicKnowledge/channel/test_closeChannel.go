package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0;i < 5;i++ {
			c <- i 
		}
		// close 可以关闭channel 停止发送数据并防止18行阻塞，同时也说明关闭channel后可以继续从channel中取数据直到channel为空
		close(c)
	}()
	/*
	for {
		// 如果不关闭则会阻塞，因为5次之后还是试图从channel中获取数据，但是此时已经没有了，就会一直等待并阻塞
		if data,ok := <-c; ok {
			fmt.Println("data = ",data)
		} else {
			break
		}
	}
	*/
	// 17-22行可以用range来代替
	for data := range c {
		fmt.Println("data = ",data)
	}
	fmt.Println("main gorountine结束")
	
}