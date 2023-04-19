package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	rsp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}