package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fileName := "./test.txt"

	file,err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		fmt.Println("openfile error = ",err)
		return
	}
	defer file.Close()

	bs := [] byte{97,98,99,100} // abcd
	first,err := file.Write(bs) // 把abcd写到文件中
	fmt.Println(first)
	HandleErr(err)
	file.WriteString("\n")

	second,err := file.WriteString("hello golang")
	fmt.Println(second)
	HandleErr(err)
}

func HandleErr(err error) {
	if err != nil{
		log.Fatal(err)
	}
}