// 策略模式
// 商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景
package main

import "fmt"

// 商场促销抽象策略
type ShopStrategy interface {
	shopping(cost float64) float64 // 根据原价得到售卖价
}

// 策略A
type AStrategy struct {

}

func (a *AStrategy) shopping(cost float64) float64{
	fmt.Println("执行策略A,所有商品打8折")
	return cost * 0.8
}

// 策略B
type BStrategy struct {

}

func (b *BStrategy) shopping(cost float64) float64{
	fmt.Println("执行策略B,所有商品满200返100")
	if cost >= 200 {
		cost -= 100
	}
	return cost
}

// 商品类
type Commodities struct {
	price float64 // 价格
	strategy ShopStrategy // 以什么策略售卖
}

// 设置策略
func (c *Commodities) setStrategy(strategy ShopStrategy) {
	c.strategy = strategy
}

func (c *Commodities) sell() {
	fmt.Println("商品的原价为:",c.price)
	fmt.Println("商品执行策略后的价格为:",c.strategy.shopping(c.price))
}

func main() {
	lining := Commodities{
		price: 200,
	}

	lining.setStrategy(new(AStrategy))
	lining.sell()

	fmt.Println("---------------")
	
	lining.setStrategy(new(BStrategy))
	lining.sell()
}