package main

import (
	"fmt"
)

func main(){
	var arr []interface{} = make([]interface{},0,8)
	fmt.Println(arr)
	fmt.Println(arr)
}

func test(v [3]int) {
	v[0] =4

	fmt.Println(v)
}