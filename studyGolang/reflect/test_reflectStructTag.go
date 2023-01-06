package main

import (
	"fmt"
	"reflect"
)

type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0;i < t.NumField();i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("taginfo:", taginfo, "tagdoc: ", tagdoc)
	}
}

func main() {
	var re Resume

	findTag(&re)
}