package infra

type Queue struct {
	q []interface{}
}

func NewQueue() *Queue {
	q := new(Queue)
	q.q = []interface{}{}
	return q
}

func (q *Queue) Enqueue(el interface{}) {
	q.q = append(q.q, el)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.q) > 0 {
		tmp := q.q[0]
		q.q = q.q[1:]
		return tmp
	}
	return nil
}

func (q *Queue) IsEmpty() bool {

	if len(q.q) == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {

	return len(q.q)
}

func (q *Queue) Peek() int {

	if len(q.q) > 0 {
		return q.q[0].(int)
	}
	return -1
}
