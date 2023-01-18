// 代理模式

package main

import "fmt"

type Goods struct {
	Kind string // 商品种类
	Fact bool // 商品真假
}

// --------抽象层--------
type Shopping interface {
	Buy(goods *Goods)
}

// ---------实现层---------
type KoreaShopping struct{

}

func (kshopping *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物，买了:",goods.Kind)
}

type JapanShopping struct{

}

func (jshopping *JapanShopping) Buy(goods *Goods) {
	fmt.Println("去日本进行了购物，买了:",goods.Kind)
}

type AmericaShopping struct{

}

func (ashopping *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物，买了:",goods.Kind)
}

// 海外代理
type ProxyShopping struct {
	shopping Shopping // 代理某个主题，可能是韩国、日本、美国代理
}

func NewProxy(shopping Shopping) Shopping {
	return &ProxyShopping{shopping: shopping}
}

func (ps *ProxyShopping) Buy(goods *Goods) {
	// + 先辨别真伪
	if (ps.distinguish(goods) == true) {
		// 调用对应的Buy()方法
		ps.shopping.Buy(goods)
		// + 海关安检
		ps.cheak(goods)
	} else {
		// 假货直接就返回了
	}

}

// 辨别某件商品的真伪
func (ps *ProxyShopping) distinguish(goods *Goods) bool {
	fmt.Println("对[" + goods.Kind + "] 进行真伪辨别")
	if goods.Fact == false {
		fmt.Printf("将要购买的商品:%s 是假货，不要购买\n",goods.Kind)
	} else {
		fmt.Printf("将要购买的商品:%s 是正品，可以购买\n",goods.Kind)
	}

	return goods.Fact
}

func (ps *ProxyShopping)cheak(goods *Goods) {
	fmt.Println("对[" + goods.Kind + "] 进行海关安检，带回祖国！")
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "哈佛大学录取通知书",
		Fact: false,
	}

	// 韩国代理
	var kshopping Shopping
	kshopping = new(KoreaShopping)

	var kproxy Shopping
	kproxy = NewProxy(kshopping)
	kproxy.Buy(&g1)

	// 美国代理
	var ashopping Shopping
	ashopping = new(AmericaShopping)

	var aproxy Shopping
	aproxy = NewProxy(ashopping)
	aproxy.Buy(&g2)

}