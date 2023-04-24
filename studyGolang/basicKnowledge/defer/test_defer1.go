package main

import (
	"log"
	"sync"
)

func main() {
	// 1.defer的作用域只在本函数内，这就是为什么把defer放在func匿名函数就能先打印defer
    // func() {
    //      defer log.Println("defer.klayhu.")
    // }()
    
	// 2.加一个go会出现什么呢？
	// go func() {
	// 	defer log.Println("defer.klayhu.")
	// }()

	// 3.单纯的在匿名函数之前加一个go 并不能保证打印defer 因为整个程序可能在打印之前就已经退出了 可以用sync的wait来保证调用
	var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer log.Println("defer.klayhu.")
        wg.Done()
    }()
    wg.Wait()
	log.Println("main.klayhu.")
}