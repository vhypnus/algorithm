package main

import (
	"fmt"
)

type Sample struct {
	val int
}


func main(){
	var s = "你好世界hellworld"
	fmt.Println(len(s))

	for _,v := range s {
		fmt.Println(string(v))
	}

	for _,v := range []byte(s) {
		fmt.Println(string(v))
	}
} 


func test(callback func(s string)) {

	if callback != nil {
		callback("hello")	
	}
	
}
