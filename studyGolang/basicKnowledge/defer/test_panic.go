package main

import (
	"log"
	"sync"
)

func main() {
	// panic + recover 捕获异常
	defer func() {
        if err := recover(); err != nil {
            log.Printf("recover: %v", err)
        }
    }()


	// 去掉defer也不行
	// if err := recover(); err != nil {
	// 	log.Printf("recover : %v", err)
	// }

	// 如果在外层再加一个gorountine就会报错
	// go func ()  {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			log.Printf("recover: %v", err)
	// 		}
	// 	}()
	// }()

	panic("klayhu.")

}