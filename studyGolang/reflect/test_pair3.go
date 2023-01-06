package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {

}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
} 

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}


func main() {
	// b: pair<type:Book, value:Book{}地址
	b := &Book{}

	// r: pair<type:nil,value:nil>
	var r Reader
	// r: pair<type:Book, value:Book{}地址
	r = b

	r.ReadBook()

	var w Writer

	w = b

	w.WriteBook()
}