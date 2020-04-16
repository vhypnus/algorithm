/*
A -> []int
n -> number
i -> index
p -> parent 
l -> left
r -> right

堆的性质：
1. l,r = 2*p,2*p+1
2. 
	1).最大堆：A[p] >= A[l] && A[P] >= A[r]
	2).最小堆: A[P] <= A[l] && A[p] <= A[r]
3.堆排序的步骤
	1).构建最大堆
	2).排序
*/

package sort

import (
	// "fmt"

)


func maxHeapify(data []int,i int) {

	n,p,l,r := len(data),i,i *2 ,i*2 +1 

	if l < n && data[l] > data[p] {
		p = l
	}

	if r < n && data[r] > data[p] {
		p = r
	}

	if p != i {
		swap(data,i,p)
		maxHeapify(data,p)
	}
}




func buildMaxHeap(data []int){
	n := len(data)
	for i:=n/2 ; i>= 0 ;i-- {
		maxHeapify(data,i)
	}
}

func HeapSort(data []int){

	buildMaxHeap(data)

	n := len(data)
	for i := n - 2 ;i >= 0 ;i -- {
		swap(data,0,i+1)

		maxHeapify(data[0:i],0)	
	}
}

func swap (data []int,s int,t int){
	tmp := data[s]
	data[s] = data[t]
	data[t] = tmp
}





