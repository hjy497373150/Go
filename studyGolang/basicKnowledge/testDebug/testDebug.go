package main

import (
	"fmt"
	"log"
)

// var myvar = 1 // is ok
// myvar := 1 //error

func main() {
	fmt.Println("testDebug1....")
	fmt.Println("testDebug2....")
	fmt.Println("testDebug3....")
	str := "hello"
	for _,ch := range []rune(str){
		// fmt.Printf("%c",ch)
		fmt.Println(rune(ch))
	}
	y := []int{
		3,
		4,
	}
	fmt.Println(y)
	log.Fatalln("Fatal Level: log entry")
}