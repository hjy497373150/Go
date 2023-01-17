package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		1.func NewTimer(d Duration) *Timer
			创建一个计时器：d时间以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
	 */
	// timer := time.NewTimer(3 *time.Second) // 创建一个3s的定时器
	// fmt.Printf("%T\n", timer)
	// fmt.Println(time.Now()) // 打印当前时间

	// // 等待channel中的信号，执行这段代码会阻塞3s
	// ch2 := timer.C
	// fmt.Println(<-ch2)

	/*
		2.func (t *timer) Stop() bool
		停止定时器，不会触发
	*/
	// fmt.Println("-----------------------")

	// // 新建一个定时器 5s后触发
	// timer2 := time.NewTimer(5 * time.Second) 

	// go func() {
	// 	// 等触发的信号
	// 	<-timer2.C
	// 	fmt.Println("timer2 结束")
	// }()
	// fmt.Println(time.Now())
	// time.Sleep(3 * time.Second) // 3s < 5s触发之前把timer停掉,如果＞5s则timer触发，stop就没意义了
	// stop := timer2.Stop()

	// if stop {
	// 	fmt.Println(time.Now())
	// 	fmt.Println("timer2 停止")
	// }

	/*
		func After(d Duration) <-chan Time
			返回一个通道：chan，存储的是d时间间隔后的当前时间。
	 */
	 ch1 := time.After(3 * time.Second) //3s后
	 fmt.Printf("%T\n", ch1) // <-chan time.Time
	 fmt.Println(time.Now()) 
	 time2 := <-ch1
	 fmt.Println(time2)

}