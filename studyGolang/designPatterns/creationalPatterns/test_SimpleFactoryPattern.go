// 简单工厂模式

package main

import "fmt"

// ------抽象层
type Fruit interface {
	Show()
}

// -------实现层
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("这里是苹果apple")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("这里是梨子pear")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("这里是香蕉banana")
}

// -----工厂模块（新增）
type Factory struct {

}

// 返回抽象产品类，apple pear banana都继承自抽象类Fruit
func (factory *Factory)CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		// apple构造初始化业务
		fruit = new(Apple) //满足多态条件的赋值，父类指向子类对象
	} else if kind == "pear" {
		fruit = new(Pear)
	} else if kind == "banana" {
		fruit = new(Banana)
	}

	return fruit
}


// ----------业务逻辑层
func main() {
	factory := new(Factory)

	// 业务逻辑层面向抽象层开发，不关心具体的类
	apple := factory.CreateFruit("apple")
	apple.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()
}
