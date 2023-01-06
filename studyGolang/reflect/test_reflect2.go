package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int 
	Name string
	Age int
}

func (this *User) Call() {
	fmt.Println("User is called...")
	fmt.Printf("%v\n",this)
}

func DoFiledAndMethod(input interface{}) {
	// 获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is : ", inputType.Name())

	// 获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is : ", inputValue)

	// 通过type 获取input中的字段
	// 1.获取interface中的reflect.type,通过type得到numfield，然后遍历
	// 2.得到每个field的数据类型
	// 3.通过field有一个interface()方法得到value
	for i := 0;i < inputType.NumField();i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s : %v = %v\n", field.Name, field.Type, value)
	}

	// 通过type 获取input中的方法，调用
	for i := 0;i < inputType.NumMethod();i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n",m.Name, m.Type)
	}
}

func main() {
	user := User{1, "klayhu", 24}

	DoFiledAndMethod(user)
}