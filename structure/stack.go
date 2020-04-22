package structure

import (

)

type Stack struct {
	
	val []int 


}

func (s *Stack) Push(v int) {
	if s.val == nil {
		//性能较差，因为pop会导致内存copy
		s.val = make([]int,0,8)
	}
	s.val = append(s.val,v)
}


func (s *Stack) Pop() int {
	if len(s.val) <= 0 {
		//panic
	}
	var r = s.val[len(s.val)-1]
	s.val = append(s.val[:len(s.val)-1],s.val[len(s.val):]...)
	return r
}
