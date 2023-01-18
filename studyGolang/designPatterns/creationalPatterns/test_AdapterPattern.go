// 适配器模式

package main

import "fmt"

// 适配的目标
type V5 interface {
	Use5V()
}

// 被适配的角色，适配者
type V220 struct{

}
func (v220 *V220)Use220V() {
	fmt.Println("使用220V电压")
}

// 适配器(实现适配的目标),是一个V5的类型
type Adapter struct {
	v220 *V220
}

func(a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")

	// 调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

// 业务类，依赖V5接口
type Phone struct {
	v5 V5
}

func NewPhone(v5 V5) *Phone {
	return &Phone{v5: v5}
}

func (p *Phone) Charge() {
	fmt.Println("Phone进行充电")
	p.v5.Use5V()
}


func main() {
	iphone := NewPhone(NewAdapter(new(V220))) // NewAdapter用V220的类型生成了一个V5的类型
	iphone.Charge()
}