// 不符合开闭原则，新增功能，需要修改已有的代码
package main

import "fmt"

type Banker struct {

}

// 存款业务
func (b *Banker) Save() {
	fmt.Println("这里是存款业务")
}

// 支付业务
func (b *Banker) Pay() {
	fmt.Println("这里是支付业务")
}

// 股票业务
func (b *Banker) Shares() {
	fmt.Println("这里是股票业务")
}

// 转账业务 (+)
func (b *Banker) Transfer() {
	fmt.Println("这里是转账业务")
	// .... 可能会影响到其他业务
}

func main() {
	banker := &Banker{}
	banker.Save()
	banker.Pay()
	banker.Shares()
	banker.Transfer()
}