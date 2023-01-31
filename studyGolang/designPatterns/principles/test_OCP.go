//开闭原则：类的改动是通过增加代码进行的，而不是修改源代码

package main

import "fmt"

// 抽象的接口
type AbstractBanker interface {
	Dobusiness()
}

// 具体的类来实现抽象的接口
type PayBanker struct {
	// 实现AbstractBanker
}

func (pb *PayBanker) Dobusiness() {
	fmt.Println("这里是支付业务")
}

type SaveBanker struct {
	// 实现AbstractBanker
}

func (sb *SaveBanker) Dobusiness() {
	fmt.Println("这里是存款业务")
}

type SharesBanker struct {
	// 实现AbstractBanker
}

func (sb *SharesBanker) Dobusiness() {
	fmt.Println("这里是股票业务")
}

// (+) 转账业务，新增一个类
type TransferBanker struct {

}

func (tb *TransferBanker) Dobusiness() {
	fmt.Println("这里是转账业务")
}

// 实现一个架构层，参数是interface
func BankBusiness(banker AbstractBanker) {
	banker.Dobusiness()
}

func main() {
	/*
	// 支付的业务
	pb := PayBanker{}
	pb.Dobusiness()

	// 存款的业务
	sb1 := SaveBanker{}
	sb1.Dobusiness()

	// 股票的业务
	sb2 := SharesBanker{}
	sb2.Dobusiness()
	*/

	BankBusiness(&PayBanker{})
	BankBusiness(&SaveBanker{})
	BankBusiness(&SharesBanker{})
}