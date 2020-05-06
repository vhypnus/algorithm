package geek

import (
	"testing"
	"log"
)


func TestRotation(t *testing.T) {
	var ar = []int{1, 2, 3, 4, 5, 6, 7}
	Rotation(ar,3)
	log.Print(ar)

	ar = []int{1, 2, 3, 4, 5, 6}
	Rotation(ar,2)
	log.Print(ar)

	ar = []int{1, 2, 3, 4, 5, 6}
	Rotation(ar,7)
	log.Print(ar)
}

func TestLargest(t *testing.T) {
	var ar = []int{3, 4, 5,6,7,8, 1, 2}

	log.Printf("result %v",Largest(ar,0,len(ar)-1))

	ar = []int{8, 1, 2,3, 4, 5,6,7}
	log.Printf("result %v",Largest(ar,0,len(ar)-1))
	ar = []int{1, 2,3, 4, 5,6,7}
	log.Printf("result %v",Largest(ar,0,len(ar)-1))
}

func TestBinary_search(t *testing.T) {
	var ar = []int{3, 4, 5,6,7,8, 1, 2}
	log.Print(ar)
	var s,e = 6,5

	log.Printf("result %v",Binary_search(ar,7,s,e))
	log.Printf("result % v",Binary_search(ar,1,s,e))
	log.Printf("result % v",Binary_search(ar,2,s,e))
	log.Printf("result % v",Binary_search(ar,3,s,e))
}

func TestMiddle(t *testing.T) {
	log.Print(middle(7,6,2))
}


func TestSwap(t *testing.T) {
	var ar = []int{1,9,3,4,5,2}
	Swap(ar,2,4)
	log.Print(ar)
}