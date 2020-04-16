package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var l1 = ArrayToList([]int{1})
	var l2 = ArrayToList([]int{9,9,9})

	var result = addTwoNumbers(l1,l2)
	for result != nil {
		fmt.Println(result)

		//reset 
		result = result.Next
	}


	l1 = ArrayToList([]int{2,4,3})
	l2 = ArrayToList([]int{5,6,4})

	result = addTwoNumbers(l1,l2)
	for result != nil {
		fmt.Println(result)

		//reset 
		result = result.Next
	}
	
}