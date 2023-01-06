package main

import "fmt"

func main() {
	var a string
	// 每一个变量内部都是一个pair
	// pair<type:string,value:abcde>
	a = "abcde"

	var allType interface{}
	allType = a

	str, _ := allType.(string)

	fmt.Println(str)

}