// 合成复用原则，Composite Reuse Principle，简称CRP。对于继承和组合，优先使用组合。
// 需求：给小猫添加一个睡觉的功能方法


package main

import "fmt"

type Cat struct {

}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 使用继承的方式来实现
type CatB struct {
	Cat // 继承Cat的所有方法和属性
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 使用组合的方式来实现
type CatC struct {
	// C *Cat
}

func (cc *CatC) Eat(cat *Cat) {
	// cc.C.Eat()
	cat.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

func main() {
	c := &Cat{}
	c.Eat()

	fmt.Println("----------")
	cb := &CatB{}
	cb.Eat()
	cb.Sleep()

	fmt.Println("----------")
	cc := &CatC{}
	cc.Eat(c)
	cc.Sleep()
}