// 饿汉实现

package main

import "fmt"

/*
	保证一个类只能有一个实例对象，而且这个实例对象还能被系统的其他模块使用
*/

// 1.保证这个类非公有化，那么外界不能通过这个类直接创建对象
// 根据golang特性，类名称首字母要小写
type singleton struct {

}

// 2.还要有一个指针可以始终指向该对象，且指针的方向不能被改变
// golang没有常指针的概念，同样也需要使该指针私有化从而不能被外界访问
var instance *singleton = new(singleton)

// 3.如果全部是私有化，那么外部模块将永远访问不到该对象
// 因此需要提供一个方法使外界能够访问
// GetInstance能否定义成一个成员方法呢？答案是否定的，因为如果是一个成员方法，那么首先需要先访问对象再访问函数
// 但是类和对象都已经私有化，外界不能访问，所以这个方法只能是一个全局普通函数
func GetInstance() *singleton {
	return instance
}

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