// 命令模式练习
// 联想路边撸串烧烤场景，有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员MM。根据命令模式，设计烤串场景。

package main

import "fmt"

// 烤串师傅，命令的接收者
type BBQMaster struct {

}

// 烤羊肉，烧鸡翅命令
func (m *BBQMaster) RoastLamb() {
	fmt.Println("烤串师傅烤羊肉")
}

func (m *BBQMaster) RoastChicken() {
	fmt.Println("烤串师傅烤鸡翅")
}

// 抽象的命令
type RoastCommand interface {
	Roast()
}

// 实现RoastCommand接口，并在实现中调用对应的烤串师傅的命令
type RoastLambCommand struct {
	master *BBQMaster
}

func (rlc *RoastLambCommand) Roast() {
	rlc.master.RoastLamb()
}

type RoastChickenCommand struct {
	master *BBQMaster
}

func (rcc *RoastChickenCommand) Roast() {
	rcc.master.RoastChicken()
}

// 服务员MM
type Waiter struct {
	cmdList []RoastCommand
}

func (w *Waiter) ServeFood() {
	if w.cmdList == nil {
		return 
	}

	for _,cmd := range w.cmdList {
		cmd.Roast() // 多态
	}
}

func main() {
	bbqmaster := new(BBQMaster)

	roastLambCmd := RoastLambCommand{master: bbqmaster}
	roastChickenCmd := RoastChickenCommand{master: bbqmaster}

	waiter := new(Waiter)
	waiter.cmdList = append(waiter.cmdList, &roastLambCmd, &roastChickenCmd)

	waiter.ServeFood()
}