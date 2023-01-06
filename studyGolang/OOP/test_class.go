package main

import "fmt"

// 类名首字母大写说明其他包也能够访问
type Hero struct {
	// 类的属性名大写，表示该属性能够对外访问，否则只能类的内部访问
	Name string
	Ad int
	Level int
}
// 值传递 copy
/*
func (this Hero) Show () {
	fmt.Println("name = ",this.name)
	fmt.Println("ad = ",this.ad)
	fmt.Println("level = ",this.level)
}

func (this Hero) GetName() {
	fmt.Println("name is ",this.name)
}

func (this Hero) SetName(newName string) {
	this.name = newName
	
}
*/

// 指针传递
func (this *Hero) Show () {
	fmt.Println("name = ",this.Name)
	fmt.Println("ad = ",this.Ad)
	fmt.Println("level = ",this.Level)
}

func (this *Hero) GetName() {
	fmt.Println("name is ",this.Name)
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
	
}
func main()  {
	hero := Hero{Name : "hjy", Ad : 100, Level : 1}
	hero.Show()
	hero.SetName("123")
	hero.Show()
}