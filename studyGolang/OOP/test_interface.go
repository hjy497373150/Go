package main

import "fmt"

// 本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

// 具体的类
type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog is sleep")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func ShowAnimal(animal Animal) {
	animal.Sleep() // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	// var animal Animal // 接口的数据类型
	// animal = &Cat{"white"}
	// animal.Sleep()

	// animal = &Dog{"Black"}
	// animal.Sleep()

	cat := Cat{"white"}
	ShowAnimal(&cat) // 传地址进去

	dog := Dog{"Black"}
	ShowAnimal(&dog)

}