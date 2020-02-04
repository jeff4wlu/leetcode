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
