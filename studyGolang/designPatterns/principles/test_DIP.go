// 依赖倒转原则：针对对接口编程，依赖于抽象（interface接口）而不依赖于具体。
package main

import "fmt"

// -------抽象层--------
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// -------实现层---------
type BenZ struct{

}

func (bz *BenZ)Run(){
	fmt.Println("奔驰正在运行")
}

type Bmw struct {

}

func (bmw *Bmw)Run(){
	fmt.Println("宝马正在运行")
}

type Zhangsan struct {

}

func (z3 *Zhangsan)Drive(car Car) {
	fmt.Println("zhang3 正在开车")
	car.Run()
}

type Lisi struct {

}

func (l4 *Lisi)Drive(car Car) {
	fmt.Println("Li4 正在开车")
	car.Run()
}


// (+)
type Wangwu struct {

}

func (w5 *Wangwu)Drive(car Car) {
	fmt.Println("Wang5 正在开车")
	car.Run()
}

// 业务逻辑层
func main() {
	// 张三开奔驰
	var benz Car
	benz = new(BenZ)

	var zhang3 Driver
	zhang3 = new(Zhangsan)
	zhang3.Drive(benz)

	// 李四开宝马
	var bmw Car
	bmw = new(Bmw)

	var li4 Driver
	li4 = new(Lisi)
	li4.Drive(bmw)

	// 王五开奔驰（+）

	var wang5 Driver
	wang5 = new(Wangwu)
	wang5.Drive(benz)
}