package structure


import (
	"testing"
	"fmt"
)

func TestQueue(t *testing.T){

	var q = NewQueue(8)
	q.Push(3)
	q.Push(4)
	fmt.Println(q)

	fmt.Println(q.Pop())
	fmt.Println(q)
}