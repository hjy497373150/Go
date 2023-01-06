package main

import (
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Price int `json:"rmb"`
	Actors []string	`json:"actors"`
}

func main() {
	movie := Movie{"Avatar", 2008, 50, []string{"zhang3", "li4"}}

	// struct -----> json
	jsonStr,err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	// json ---->struct

	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error")
		return
	}
	fmt.Printf("%v\n", myMovie)
}