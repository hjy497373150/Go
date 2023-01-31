// 模版方法模式

package main

import "fmt"

// 抽象类，包裹一个模版的全部实现步骤
type Beverage interface {

	BoilWater() // 煮开水
	Brew() // 冲泡
	PourInCup() // 倒入杯中
	AddThings() // 添加小料

	IfAddThings() bool //是否需要添加小料,具体流程可以重写该方法
}

// 模版基类,封装一套流程模版，让具体的制作流程继承并实现
type template struct {
	b Beverage //该属性中的方法由具体的类来实现
}

// 封装的固定模版
func (t *template) MakeBeverage() {
	if t.b == nil {
		return
	}
	t.b.BoilWater() // 留给子类去具体实现
	t.b.Brew()
	t.b.PourInCup()
	if t.b.IfAddThings() == true {
		t.b.AddThings()
	}

}

// 具体的模版子类,制作咖啡
type MakeCoffee struct {
	template // 继承模版
}

// 实现所有步骤的方法
func (mc *MakeCoffee) BoilWater() {
	fmt.Println("将水烧开到100摄氏度")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("用水冲咖啡粉")
}

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("将冲好的咖啡倒入杯中")
}

func (mc *MakeCoffee) AddThings() {
	fmt.Println("添加牛奶和糖")
}

func (mc *MakeCoffee) IfAddThings() bool {
	return true
}

// 构造函数
func NewMakeCoffee() *MakeCoffee {
	makecoffee := new(MakeCoffee)
	// b为Beverage 是makecoffee的接口 需要给接口赋值，指向具体的子类对象
	makecoffee.b = makecoffee
	return makecoffee
}

// 具体的模版子类，制作茶
type MakeTea struct {
	template
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("将水烧开到80摄氏度")
}

func (mt *MakeTea) Brew() {
	fmt.Println("用水冲茶")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("将冲好的茶倒入杯中")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) IfAddThings() bool {
	return false
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func main() {
	// 1.制作一杯咖啡
	makecoffee := NewMakeCoffee()
	makecoffee.template.MakeBeverage()

	fmt.Println("-----------------")

	// 制作一杯茶
	makeTea := NewMakeTea()
	makeTea.template.MakeBeverage()
}