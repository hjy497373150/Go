// 单例模式懒汉实现

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	保证一个类只能有一个实例对象，而且这个实例对象还能被系统的其他模块使用
*/

type singleton struct {

}

var instance *singleton // 懒汉实现需要在使用到这个单例的时候才去创建它

var lock sync.Mutex // 定义一个锁

var initialized uint32 // 标记

func GetInstance() *singleton {
	// 标记位为1，说明单例已经被生成
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	// 如果没有，就加锁申请
	lock.Lock()

	defer lock.Unlock()
	// 只有该方法首次被调用才会生成这个单例的对象
	if instance == nil {
		instance = new(singleton)
		atomic.StoreUint32(&initialized, 1) // 设置标记位为1
		return instance
	}
	return instance
}

/*
func GetInstance() *singleton {
	// 为了线程安全，需要加锁
	lock.Lock()

	defer lock.Unlock()

	// 只有该方法首次被调用才会生成这个单例的对象
	if instance == nil {
		instance = new(singleton)
		return instance
	}
	return instance
}
*/

func (s *singleton) SomeThing() {
	fmt.Println("单例的某个方法")
}


func main() {
	s1 := GetInstance()
	s1.SomeThing()
	
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}