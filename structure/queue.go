package structure

type Queue struct {

	val []interface{}

	start int

	end int 
}

func NewQueue(capacity int) *Queue{
	var q = new(Queue)
	q.val = make([]interface{},0,capacity)
	return q 
}


func (q *Queue) Push(v interface{}) {
	q.val = append(q.val,v)
}


func (q *Queue) Pop() interface{} {
	var  v = q.val[0]
	q.val = q.val[1:len(q.val)]
	return v 
}