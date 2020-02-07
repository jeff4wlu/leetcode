package leet225

import "leetcode/infra"

//为了简单起见，本题只用适用于int

type stack struct {
	q *infra.Queue
}

func newStack() *stack {
	s := new(stack)
	s.q = infra.NewQueue()
	return s
}

func (s *stack) push(num int) {
	s.q.Enqueue(num)
}

func (s *stack) pop() int {
	n := s.q.Size()
	if n > 0 {
		for count := n - 1; count > 0; count-- {
			tmp := s.q.Dequeue().(int)
			s.q.Enqueue(tmp)
		}
		return s.q.Dequeue().(int)
	}
	return -1
}

func (s *stack) top() int {
	n := s.q.Size()
	if n > 0 {

		for count := n - 1; count > 0; count-- {
			tmp := s.q.Dequeue().(int)
			s.q.Enqueue(tmp)

		}
		res := s.q.Dequeue().(int)
		s.q.Enqueue(res)
		return res
	}
	return -1
}

func (s *stack) empty() bool {
	return s.q.IsEmpty()
}
