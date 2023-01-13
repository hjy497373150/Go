package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// trace编程的基本过程
// 1.创建trace文件
// 2.启动
// 3.停止

// 使用go tool trace trace.out来执行
func main() {
	// 1.创建一个trace文件
	f,err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 2.启动这个文件
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	
	// 正常要调试的业务代码
	fmt.Println("hello GMP!")

	// 3.停止
	trace.Stop()

}