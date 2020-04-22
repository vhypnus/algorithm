package structure


type Heap struct {
	val []int

	//0:maxHeap or 1:minHeap
	type byte 
}


func NewHeap(capacity int,type byte) *Heap{
	var h = new(Heap)
	h.val = make([]int,0,capacity)
	h.type = type
	return h 
}


func (h *Heap) Push(v int) {
	h.val = append(h.val,v)
	//reorder
}

// func 

func Pop() int {

	return 0
}