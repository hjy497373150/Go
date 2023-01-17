// 工厂方法模式

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


// -----工厂模块
// 工厂类，抽象的接口
type AbstractFactory interface {
	CreateFruit() Fruit //返回一个Fruit对象
}

// 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (appleFactory *AppleFactory)CreateFruit() Fruit {
	var apple Fruit

	// 生产一个具体的苹果
	apple = new(Apple)

	// 返回的一定是苹果
	return apple
}

// 具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (pearFactory *PearFactory)CreateFruit() Fruit {
	var pear Fruit

	// 生产一个具体的梨
	pear = new(Pear)

	// 返回的一定是梨
	return pear

}

// 具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (bananaFactory *BananaFactory)CreateFruit() Fruit {
	var banana Fruit

	// 生产一个具体的香蕉
	banana = new(Banana)

	// 返回的一定是香蕉
	return banana

}

// ----------业务逻辑层
func main () {
	// 需求1：需要一个具体的苹果对象
	// 1.先要一个苹果的工厂
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)

	// 2.生产一个具体的苹果
	var apple Fruit
	apple = appleFactory.CreateFruit()

	apple.Show() //多态

	
	var pearFactory AbstractFactory
	pearFactory = new(PearFactory)

	
	var pear Fruit
	pear = pearFactory.CreateFruit()

	pear.Show() //多态

	
	var bananaFactory AbstractFactory
	bananaFactory = new(BananaFactory)

	
	var banana Fruit
	banana = bananaFactory.CreateFruit()

	banana.Show() //多态

}