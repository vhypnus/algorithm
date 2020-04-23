package structure

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	var h = NewHeap(8,0)

	fmt.Println(h)

	h.Push(1)
	h.Push(2)
	h.Push(3)
	h.Push(4)
	h.Push(5) 
	fmt.Println(h)

	fmt.Println(h.Pop())
	fmt.Println(h)
}