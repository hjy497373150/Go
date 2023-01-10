package main

import "fmt"

type Human struct {
	Name string 
	Sex string
}

func (this *Human) Work() {
	fmt.Println("Human work...")
}

func (this *Human) Eat() {
	fmt.Println("Human eat...")
}

type SuperMan struct {
	Human // 子类继承，包括类的属性和方法
	Level int
}

// 子类重写方法
func (this *SuperMan) Work() {
	fmt.Println("Superman work...")
}

// 子类的新方法
func (this *SuperMan) fly() {
	fmt.Println("Superman fly...")
}

func (this *SuperMan) Print() {
	fmt.Println("name = " ,this.Name)
	fmt.Println("sex = " ,this.Sex)
	fmt.Println("level = " ,this.Level)
}

func main() {
	h := Human{"zhang3", "male"}
	h.Work()
	h.Eat()

	// 定义子类对象
	// s := SuperMan{Human{"l14", "female"}, 100}
	var s SuperMan
	s.Name = "li4"
	s.Sex = "female"
	s.Level = 100

	s.Work()
	s.Eat() // 调用的是父类的方法
	s.fly()

	s.Print()
}