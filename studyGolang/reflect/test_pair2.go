package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// klayhu:pair<type:*os.File, value:"/dev/tty" 文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	// r: pair<type :nil, value: nil>
	var r io.Reader

	// r:pair<type:*os.File, value:"/dev/tty" 文件描述符>
	r = tty

	var w io.Writer
	// r:pair<type:*os.File, value:"/dev/tty" 文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("this is a test...\n"))
}