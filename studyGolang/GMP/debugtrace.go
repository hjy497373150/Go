package main

import (
	"fmt"
	"runtime"
	"time"
)

// GODEBUG=schedtrace=1000 ./debugtrace 运行 1000指1000ms打印一次
// SCHED 调试的信息
// 0ms 从程序启动到输出经历的时间
// gomaxprocs P的数量 一般默认和CPU核心数一样
// idleprocs 处理idle（空闲）状态的P的数量 gomaxprocs-idleprocs = 目前正在执行的P的数量
// threads 线程数量（包括M0,包括GODEBUG调试的线程）
// spinningthreads 处于自旋状态的thread数量
// idlethreads 处于idle状态的thread的数量
// runqueue 全局队列中的G数量
// [0,0,...] 每个P的本地队列中，目前存在的G的数量
func main() {
	fmt.Println("CPU核数:",runtime.NumCPU())
	for i := 0;i<5;i++ {
		time.Sleep(time.Second)
		fmt.Println("hello GMP...")
	}
}