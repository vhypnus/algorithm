package structure

import(
	// "fmt"
)

type Heap struct {
	val []int

	//0:maxHeap or 1:minHeap
	t byte 
}


func NewHeap(capacity int,t byte) *Heap{
	var h = new(Heap)
	h.val = make([]int,0,capacity)
	h.t = t
	return h 
}


func (h *Heap) Push(v int) {
	h.val = append(h.val,v)
	var hv = h.val

	for max,p := len(hv)-1, 0 ;max > 0 ; max = p {

		if max << 1 & 2 == 0 {
			p = (max-2) >> 1
		} else {
			p = (max-1) >> 1
		}


		if hv[max] > hv[p] {
			tmp := hv[p]
			hv[p] = hv[max]
			hv[max] = tmp
		}

	}
}


func (h *Heap) Pop() int {

	var v = h.val[0]
	h.val = h.val[1:]

	var hv = h.val

	for max,p := len(hv)-1, 0 ;max > 0 ; max = p {

		if max << 1 & 2 == 0 {
			p = (max-2) >> 1
		} else {
			p = (max-1) >> 1
		}


		if hv[max] > hv[p] {
			tmp := hv[p]
			hv[p] = hv[max]
			hv[max] = tmp
		}

	}

	return v
}