package main

import (
	"fmt"
)

type Sample struct {
	val int
}

type Student struct {

}

func main(){
	var s1 = &Sample{17}
	var s2 = &Sample{18}
	var s3 = s2

	var m = make(map[*Sample]*Sample)
	m[s1] = s1
	m[s2] = s2
	fmt.Println(s2 == s3)
	fmt.Println(m[s3])
}