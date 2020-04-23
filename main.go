package main

import (
	"fmt"
)

func main(){
	fmt.Println( 2 >> 1)
}

func test(v [3]int) {
	v[0] =4

	fmt.Println(v)
}