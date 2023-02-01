//普通切片
package main

import "fmt"

func main() {
	// 下面两种定义方式都可以
	var numbers []int = make([]int,3,5)
//    numbers := make([]int,3,5)

   printSlice(numbers)
}


func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}