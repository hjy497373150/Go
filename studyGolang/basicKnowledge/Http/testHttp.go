package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("hello, now is %s\n", time.Now().String()))
	})
	if err := http.ListenAndServe(":8088", nil);err != nil {
		fmt.Printf("listen error: %s\n", err.Error())
	}
}