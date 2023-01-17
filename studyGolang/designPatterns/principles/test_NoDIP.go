// 耦合度极高的模块设计
// 实现张三开奔驰，李四开宝马的需求
// 如果要实现张三开宝马，李四开奔驰，就会使代码修改极多，且如果业务比较复杂，修改极有可能会带来额外的问题
package main

import "fmt"

// 宝马类
type BMW struct {

}

func (bmw *BMW)Run() {
	fmt.Println("BMW is running...")
}

// 奔驰类
type Benz struct {

}

func (benz *Benz)Run() {
	fmt.Println("Benz is running...")
}

// 张三类
type Zhang3 struct{

}

func (z3 *Zhang3)DriveBenz(benz *Benz) {
	fmt.Println("zhang3 drive Benz...")
	benz.Run()
}

// 李四类
type Li4 struct{

}

func (l4 *Li4)DriveBMW(bmw *BMW) {
	fmt.Println("li4 drive BMW...")
	bmw.Run()
}


func main() {
	benz := Benz{}
	zhang3 := Zhang3{}
	zhang3.DriveBenz(&benz)

	bmw := BMW{}
	li4 := Li4{}
	li4.DriveBMW(&bmw)
}