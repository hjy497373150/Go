package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 利用read和write来边读边写
func copyFile1(srcFile,destFile string)(int,error) {
	file1,err := os.Open(srcFile)
	if err != nil {
		fmt.Println("open srcFile err:",err)
		return 0,err
	}

	file2,err := os.Create(destFile) // 创建destFile ，不存在就创建，存在就清空原来的文件内容
	if err != nil {
		fmt.Println("open destFile err:",err)
		return 0,err
	}
	// 结束后关闭文件
	defer file1.Close()
	defer file2.Close()

	// 开始拷贝
	bs := make([]byte,1024,1024) // 每次读1024个字节
	n := -1 //读取的数据量
	total := 0 // 拷贝的总字节数
	for {
		n,err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕！")
			break
		} else if err != nil {
			fmt.Println("拷贝报错！")
			return total,err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total,nil
}

// 直接利用copy()库函数来实现
func copyFile2(srcFile,destFile string)(int64,error) {
	file1,err := os.Open(srcFile)
	if err != nil {
		fmt.Println("open srcFile err:",err)
		return 0,err
	}

	file2,err := os.Create(destFile) // 创建destFile ，不存在就创建，存在就清空原来的文件内容
	if err != nil {
		fmt.Println("open destFile err:",err)
		return 0,err
	}
	// 结束后关闭文件
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2,file1) // 把file1的内容copy到file2
}

// 使用ioutil.readfile()和ioutil.writefile() 这种方式是一次读取和一次性写入，不太适合大文件，容易内存溢出
func copyFile3(srcFile,destFile string)(int,error) {
	input,err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println("read srcFile err:",err)
		return 0,err
	}

	err = ioutil.WriteFile(destFile,input,0644)
	if err != nil {
		fmt.Println("write to destFile err:",err)
		return 0,err
	}

	return len(input),nil
}


func main() {
	srcFile := "./test.txt"
	destFile := "./testDest.txt"
	total,err := copyFile2(srcFile, destFile)
	fmt.Println(err)
	fmt.Println(total)
}