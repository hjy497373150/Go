package main

import "fmt"

type Person struct {
	name string
}

func (this Person) Show() {
	fmt.Println("Person--->",this.name)
}

// 类型别名
type People = Person

func (this People) Show2() {
	fmt.Println("People--->",this.name)
}

type Student struct {
	// 嵌入两个结构
	Person
	People
}

func main() {
	var s Student

	// s.name = "zhang3" // 报错 ambiguous selector s.name
	s.People.name = "zhang3"
	s.Person.name = "li4"

	// s.Show() // 报错 ambiguous selector s.Show
	s.Person.Show()
	s.People.Show2()
	fmt.Printf("%T, %T\n",s.Person, s.People) // main.Person, main.Person
}