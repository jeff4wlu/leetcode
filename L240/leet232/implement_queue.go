package leet232

import "leetcode/infra"

type queue struct {
	sin  infra.Stack
	sout infra.Stack
}

func newQueue() *queue {
	q := new(queue)
	q.sin = infra.CreateStack()
	q.sout = infra.CreateStack()
	return q
}

func (q *queue) push(num int) {
	q.sin.Push(num)
}

//为了简单-1代表没有元素
func (q *queue) pop() int {
	if q.sout.IsEmpty() {
		for !q.sin.IsEmpty() {
			q.sout.Push(q.sin.Pop().(int))
		}
		if q.sout.IsEmpty() {
			return -1
		}
	}
	return q.sout.Pop().(int)
}

func (q *queue) peek() int {
	if q.sout.IsEmpty() {
		for !q.sin.IsEmpty() {
			q.sout.Push(q.sin.Pop().(int))
		}
		if q.sout.IsEmpty() {
			return -1
		}
	}
	tmp := q.sout.Pop().(int)
	q.sout.Push(tmp)
	return tmp

}

func (q *queue) empty() bool {
	return q.sin.IsEmpty() && q.sout.IsEmpty()
}
