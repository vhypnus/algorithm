package structure

import (

)

type Stack struct {
	
	val []int 

	len int 

	start int

	end int


}

func (s *Stack) Push(v int) {
	if s.val == nil {
		s.val = make([]int,0,8)
	}
	s.val = append(s.val,v)
}


func (s *Stack) Pop() int {
	if len(s.val) <= 0 {
		//panic
	}

	var r = s.val[len(s.val)-1]
	//这会有性能问题
	s.val = append(s.val[:len(s.val)-1],s.val[len(s.val):]...)
	return r
}
