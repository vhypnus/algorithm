package structure

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T){

	var s = new(Stack)

	s.Push(1)
	s.Push(2)

	fmt.Println(s)

	fmt.Println(s.Pop())
	s.Push(3)
	s.Push(4)


	fmt.Println(s)
	fmt.Println(s.Pop())
	s.Pop()
	fmt.Println(s)


	var arr = make([]int,0,8)
	arr = append(arr,1)
	arr = append(arr,2)
	arr = append(arr,3)
	arr = append(arr,4)
	arr = append(arr,5)
	fmt.Println(cap(arr))
	fmt.Println(len(arr))

	arr = append(arr[0:2],arr[3:]...)
	fmt.Println(arr)
	fmt.Println(cap(arr))
	fmt.Println(len(arr))
}