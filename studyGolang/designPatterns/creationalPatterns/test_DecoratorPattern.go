// 装饰器模式

package main

import "fmt"

// -----------抽象层
// 抽象的Phone类
type Phone interface {
	Show()
}

// 抽象的装饰器基础类，本应该是interface，但是golang interface不能有成员属性，所以只能是struct
type Decorator struct {
	phone Phone
}

func (d *Decorator)Show() {

}

// -------------实现层
type Huawei struct {

}

func (h *Huawei) Show() {
	fmt.Println("秀出了Huawei手机...")
}

type Iphone struct {

}

func (i *Iphone) Show() {
	fmt.Println("秀出了Iphone手机...")
}

type MoDecorator struct {
	Decorator //继承装饰器基础类，主要继承Phone属性
}

func (md *MoDecorator) Show() {
	md.phone.Show() // 调用被装饰构件的原方法
	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone: phone}}
}

type KeDecorator struct {
	Decorator //继承装饰器基础类，主要继承Phone属性
}

func (kd *KeDecorator) Show() {
	kd.phone.Show() // 调用被装饰构件的原方法
	fmt.Println("手机壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone: phone}}
}

// ------------业务逻辑层
func main() {
	var huawei Phone
	huawei = new(Huawei)
	huawei.Show()

	fmt.Println("---------------")
	// 都是面向抽象层写业务
	var mohuawei Phone
	mohuawei = NewMoDecorator(huawei) //通过huawei new 一个mohuawei出来
	mohuawei.Show()

	fmt.Println("---------------")
	var kemohuawei Phone
	kemohuawei = NewKeDecorator(mohuawei) // 可以多次装饰
	kemohuawei.Show()
}