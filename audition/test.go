package main

import "fmt"


func main() {
    slice := make([]int, 2, 3)
    for i := 0; i < len(slice); i++ {
        slice[i] = i
    }

    ret := changeSlice(slice)
    ret[1] = 111

    fmt.Printf("slice: %v, ret: %v \n", slice, ret)
}

func changeSlice(s []int) []int {
    s[0] = 10
    s = append(s, 3)
    s = append(s, 4)
    s[1] = 100
    
    return s
}
