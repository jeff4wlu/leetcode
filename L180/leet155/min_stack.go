package leet155

import "leetcode/infra"

type stackCore struct {
	c []int
}

type stack struct {
	c   *stackCore
	min *stackCore
}

func NewStack() *stack {
	s := new(stack)
	s.c = new(stackCore)
	s.c.c = make([]int, 10)
	s.min = new(stackCore)
	s.min.c = make([]int, 10)
	return s
}

func (s *stack) pop() int {
	s.min.pop()
	return s.c.pop()
}

func (s *stack) push(a int) {
	s.c.push(a)
	tmp := s.min.top()
	if a < tmp {
		s.min.push(a)
	} else {
		s.min.push(tmp)
	}
}

func (s *stack) top() int {
	return s.c.top()
}

func (s *stack) getMin() int {
	return s.min.top()
}

func (s *stackCore) push(a int) {

	s.c = append(s.c, a)
}

func (s *stackCore) pop() int {
	n := len(s.c)
	if n > 0 {
		tmp := s.c[n-1]
		s.c = s.c[:n-1]
		return tmp
	}
	return infra.INT32_MAX
}

func (s *stackCore) top() int {
	n := len(s.c)
	if n > 0 {
		return s.c[n-1]
	}
	return infra.INT32_MAX
}
