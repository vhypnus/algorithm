package structure

type Queue struct {

	val []interface{}

	num int 
}

func NewQueue(capacity int) *Queue{
	var q = new(Queue)
	q.val = make([]interface{},0,capacity)
	return q 
}


func (q *Queue) Push(v interface{}) {
	q.val = append(q.val,v)
	q.num += 1
}


func (q *Queue) Pop() interface{} {
	var  v = q.val[0]
	q.val = q.val[1:len(q.val)]
	q.num -= 1
	return v 
}

func (q *Queue) Size() int {
	return q.num
}