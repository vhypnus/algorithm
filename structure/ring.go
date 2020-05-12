package structure

//
type Ring struct {

	// slice
	s []interface{}

	// head
	h int

	// tail
	t int

	//lenght
	l int

	// capacity
	c int


}


func NewRing(capacity int) {
	var s = make([]interface{},0,capacity)
	var r = &Ring{s,0,0,0,capacity}
}


func (r *Ring) Push(v interface{}){
	var i int
	if l < c {
		if t = l-1 {
			t = 0
		} 
		i = t+1
	} else {
		r.resize()
		i = l
	}

	r.s[i] = v
	r.l++
	r.c++
}

func (r *Ring) Pop() {

	
	r.l--
	r.c--
}

func (r *Ring) resize(){

	var t = &Ring{make([]interface{},0,2*r.c),0,0,0,2*r.c}
	copy(r,t)
	r = t
}


func copy(s *Ring,t *Ring) {


}