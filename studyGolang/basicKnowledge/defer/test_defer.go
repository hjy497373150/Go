// defer语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为延迟函数。
// defer作用：

// 释放占用的资源
// 捕捉处理异常
// 输出日志
package main

import "fmt"

// 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。类似栈
func Demo(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

func main() {
	Demo()
}
