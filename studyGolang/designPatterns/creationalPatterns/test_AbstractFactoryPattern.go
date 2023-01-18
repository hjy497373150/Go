// 抽象工厂方法模式

package main

import "fmt"

// ------抽象层
// 抽象的苹果（抽象产品角色
type AbstractApple interface {
	ShowApple()
}

// 抽象的梨
type AbstractPear interface {
	ShowPear()
}

// 抽象的香蕉
type AbstractBanana interface {
	ShowBanana()
}

// 抽象工厂角色
type AbstractFactory interface {
	CreateApple() AbstractApple //返回一个apple对象
	CreatePear() AbstractPear //返回一个pear对象
	CreateBanana() AbstractBanana //返回一个banana对象
}

// -------实现层
// 中国产品族
// 具体产品角色
type ChinaApple struct {
	AbstractApple
}

func (chinaApple *ChinaApple) ShowApple() {
	fmt.Println("这里是中国苹果")
}

type ChinaPear struct {
	AbstractPear
}

func (chinaPear *ChinaPear) ShowPear() {
	fmt.Println("这里是中国梨子")
}

type ChinaBanana struct {
	AbstractBanana
}

func (chinaBanana *ChinaBanana) ShowBanana() {
	fmt.Println("这里是中国香蕉")
}

// 具体的中国工厂（可以生产中国苹果，中国梨，中国香蕉）
type ChinaFactory struct {
	AbstractFactory
}

func (chinaFactory *ChinaFactory)CreateApple() AbstractApple {
	var apple AbstractApple

	// 生产一个具体的中国苹果
	apple = new(ChinaApple)

	// 返回的一定是中国苹果
	return apple
}

func (chinaFactory *ChinaFactory)CreatePear() AbstractPear {
	var pear AbstractPear

	// 生产一个具体的中国梨
	pear = new(ChinaPear)

	// 返回的一定是中国梨
	return pear
}

func (chinaFactory *ChinaFactory)CreateBanana() AbstractBanana {
	var banana AbstractBanana

	// 生产一个具体的中国香蕉
	banana = new(ChinaBanana)

	// 返回的一定是中国香蕉
	return banana
}

// 日本产品族
type JapanApple struct {
	AbstractApple
}

func (japanApple *JapanApple) ShowApple() {
	fmt.Println("这里是日本苹果")
}

type JapanPear struct {
	AbstractPear
}

func (japanPear *JapanPear) ShowPear() {
	fmt.Println("这里是日本梨子")
}

type JapanBanana struct {
	AbstractBanana
}

func (japanBanana *JapanBanana) ShowBanana() {
	fmt.Println("这里是日本香蕉")
}

// 具体的日本工厂（可以生产日本苹果，日本梨，日本香蕉）
type JapanFactory struct {
	AbstractFactory
}

func (japanFactory *JapanFactory)CreateApple() AbstractApple {
	var apple AbstractApple

	// 生产一个具体的日本苹果
	apple = new(JapanApple)

	// 返回的一定是日本苹果
	return apple
}

func (japanFactory *JapanFactory)CreatePear() AbstractPear {
	var pear AbstractPear

	// 生产一个具体的日本梨
	pear = new(JapanPear)

	// 返回的一定是日本苹果
	return pear
}

func (japanFactory *JapanFactory)CreateBanana() AbstractBanana {
	var banana AbstractBanana

	// 生产一个具体的日本香蕉
	banana = new(JapanBanana)

	// 返回的一定是日本香蕉
	return banana
}

// 美国产品族
type USAApple struct {
	AbstractApple
}

func (usaApple *USAApple) ShowApple() {
	fmt.Println("这里是美国苹果")
}

type USAPear struct {
	AbstractPear
}

func (usaPear *USAPear) ShowPear() {
	fmt.Println("这里是美国梨子")
}

type USABanana struct {
	AbstractBanana
}

func (usaBanana *USABanana) ShowBanana() {
	fmt.Println("这里是美国香蕉")
}

// 具体的美国工厂（可以生产美国苹果，美国梨，美国香蕉）
type USAFactory struct {
	AbstractFactory
}

func (usaFactory *USAFactory)CreateApple() AbstractApple {
	var apple AbstractApple

	// 生产一个具体的美国苹果
	apple = new(USAApple)

	// 返回的一定是美国苹果
	return apple
}

func (usaFactory *USAFactory)CreatePear() AbstractPear {
	var pear AbstractPear

	// 生产一个具体的美国梨
	pear = new(USAPear)

	// 返回的一定是美国梨
	return pear
}

func (usaFactory *USAFactory)CreateBanana() AbstractBanana {
	var banana AbstractBanana

	// 生产一个具体的美国香蕉
	banana = new(USABanana)

	// 返回的一定是美国香蕉
	return banana
}

// 针对产品族进行添加，符合开闭原则，比如新增一个印度工厂
// 针对产品等级结构进行添加，不符合开闭原则，比如新增一个新的葡萄水果类

// ----------业务逻辑层
func main () {
	// 需求1：需要中国的苹果，香蕉，梨
	// 1.先要一个中国工厂
	var cFac AbstractFactory
	cFac = new(ChinaFactory)

	// 2.生产苹果
	var cApple AbstractApple
	cApple = cFac.CreateApple()

	cApple.ShowApple()

	// 3.生产梨
	var cPear AbstractPear
	cPear = cFac.CreatePear()

	cPear.ShowPear()

	// 4.生产香蕉
	var cBanana AbstractBanana
	cBanana = cFac.CreateBanana()

	cBanana.ShowBanana()

	// 需求2：需要日本的苹果、梨、香蕉
	// 1.先要一个日本工厂
	var jFac AbstractFactory
	jFac = new(JapanFactory)

	// 2.生产苹果
	var jApple AbstractApple
	jApple = jFac.CreateApple()

	jApple.ShowApple()

	// 3.生产梨
	var jPear AbstractPear
	jPear = jFac.CreatePear()

	jPear.ShowPear()

	// 4.生产香蕉
	var jBanana AbstractBanana
	jBanana = jFac.CreateBanana()

	jBanana.ShowBanana()
}