package main

import (
	"fmt"
	"os"
)

func main() {
	// 返回fileinfo
	/*
	FileInfo：文件信息
		interface
			Name()，文件名
			Size()，文件大小，字节为单位
			IsDir()，是否是目录
			ModTime()，修改时间
			Mode()，权限

	 */
	fileInfo, err := os.Stat("./test.txt")
	if err != nil {
		fmt.Println("err: ",err)
		return
	}
	fmt.Printf("%T\n",fileInfo)
	fmt.Println("file name = ",fileInfo.Name())
	fmt.Println("file size = ",fileInfo.Size())
	fmt.Println("file is Dir ? ",fileInfo.IsDir())
	fmt.Println("file modtime = ", fileInfo.ModTime())
	fmt.Println("fole mode = ", fileInfo.Mode())

}