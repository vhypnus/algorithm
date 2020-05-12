package main

import (
	"fmt"
)

type Say interface {
	Hello()
}



type Sample struct {
	val int
}

func Hi(say Say) {
	say.Hello()
}

func (s *Sample) Hello(){
	fmt.Println("hello world")
}

type Parent struct {

}

func (p *Parent) Yes(){
	fmt.Println("yes............")
}

type Child struct {
	Parent
}


func main(){
	var p Parent = Parent(&Child{})
	c.Yes()

} 


func test(callback func(s string)) {

	if callback != nil {
		callback("hello")	
	}
	
}
