package main

import (
	"fmt"
)

type Sample struct {
	val int
}


func main(){
	var m = make(map[interface{}]Sample)

	m[1] = Sample{}
} 


func test(v interface{}) {
	var  s = v.(Sample)
	s.val = 20

	fmt.Println(s)
}