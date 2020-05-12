package ds


type Ring struct {

	//slice
	s []interface {}

	// head
	h int

	//capacity
	c int
}

func NewRing(capacity int) *Ring {
	var s = make([]interface,0,capacity)
	var r = & Ring{s,0,capacity}
}

func (r *Ring) Push(v interface{}){
	var i int
	
	if h = c -1 {
		h = 0
	}

	r.s[h] = v
}