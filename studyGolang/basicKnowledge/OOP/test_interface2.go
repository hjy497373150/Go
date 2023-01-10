package main

import "fmt"

// interface{}是万能类型
func myFunc (arg interface{}) {
	fmt.Println("myfunc is called...")
	fmt.Println(arg)

	// interface{}如何区分 引用的底层数据类型是什么
	// 类型断言 机制
	value,ok := arg.(string)

	if !ok  {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string type, value = ",value)
	}
}

type Book struct {
	Name string
}


func main() {
	book := Book{"Golang"}
	myFunc(book)
	myFunc(123)
	myFunc(3.14)
	myFunc("123")
}