package sort 


import (
	"testing"
	"fmt"
)

// func TestMaxHeapify(t *testing.T){
// 	var data []int = []int{2,3,4,7,1,14,7,9,81}

// 	maxHeapify(data,9,0)

// 	fmt.Println(data)
// }

func TestHeapSort(t *testing.T) {
	var data []int = []int{6,3,50,7,1,14,7,9,81}
	HeapSort(data)
	fmt.Println(data)


	data = []int{1,3,5,7,14,81}
	HeapSort(data)
	fmt.Println(data)

	data = []int{1}
	HeapSort(data)
	fmt.Println(data)
}

func TestQuicksort(t *testing.T) {
	var data []int = []int{6,3,50,7,1,14,7,9,81}
	quicksort(data,0,8)
	fmt.Println(data)


	data = []int{1,2,3,4,5,6}
	quicksort(data,0,5)
	fmt.Println(data)
}