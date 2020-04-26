package main

import (
	"fmt"
)

const (

	WHITE = iota
	
	GRAY 

	BLACK

)

func main(){
	var s = make([]int,8,8)

	fmt.Println(&s)

	s = append(s,1)
	fmt.Println(&s)
}