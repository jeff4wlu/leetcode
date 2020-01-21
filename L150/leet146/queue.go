package leet146

type ItemNode struct {
	Value int
	Next  *ItemNode
	Pre   *ItemNode
}

type ItemQueue struct {
	cap       int
	size      int
	dummyHead *ItemNode
	dummyTail *ItemNode
}

func New(cap int) *ItemQueue {
	s := new(ItemQueue)
	s.cap = cap
	s.dummyHead = new(ItemNode)
	s.dummyTail = new(ItemNode)
	s.dummyHead.Next = s.dummyTail
	s.dummyTail.Pre = s.dummyHead
	return s
}

func (s *ItemQueue) Enqueue(n int) *ItemNode {

	if s.cap > s.size {
		s.size++
		node := new(ItemNode)
		node.Value = n

		tmp := s.dummyHead.Next
		s.dummyHead.Next = node
		node.Next = tmp
		tmp.Pre = node
		node.Pre = s.dummyHead
		return node
	}

	return nil
}

func (s *ItemQueue) Dequeue() *ItemNode {

	if s.size != 0 {
		tmp := s.dummyTail.Pre
		tmp.Pre.Next = s.dummyTail
		s.dummyTail.Pre = tmp.Pre
		s.size--
		return tmp
	}
	return nil

}

func (s *ItemQueue) Remove(n *ItemNode) bool {

	if n.Next != nil && n.Pre != nil && s.size > 0 {
		s.size--
		n.Pre.Next = n.Next
		n.Next.Pre = n.Pre
		return true
	}

	return false
}

func (s *ItemQueue) Size() int {
	return s.size
}

func (s *ItemQueue) Cap() int {
	return s.cap
}
