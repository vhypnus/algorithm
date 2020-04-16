

package leetcode

import (
	"testing"
	"fmt"
)

/*
func TestArrayToList(t *testing.T ){

	var ints = []int{2,3,45,6}
	var head = ArrayToList(ints)
	fmt.Println(head)
	fmt.Println(head.Next)
	fmt.Println(head.Next.Next)
	fmt.Println(head.Next.Next)
	fmt.Println(head.Next.Next.Next)
}*/

func TestConvert(t *testing.T) {

	fmt.Println(3 & 2)
	var str = "PAYPALISHIRING"
	var r = convert(str,3)
	fmt.Println(r)


	str = "ABCDEFG"
	r = convert(str,1)
	fmt.Println(r)

	str = "ABCDEFG"
	r = convert(str,7)
	fmt.Println(r)

	str = "ABC"
	r = convert(str,2)
	fmt.Println(r)
}