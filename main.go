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
	fmt.Println(len(m))
	_ = func (s string) {
		fmt.Println(s)
	}

	// fmt.Prin
	test(nil)
} 


func test(callback func(s string)) {

	if callback != nil {
		callback("hello")	
	}
	
}
