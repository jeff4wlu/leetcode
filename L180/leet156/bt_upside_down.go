package leet156

import "leetcode/infra"
import "github.com/golang-collections/collections/stack"

func BTUpsideDown(tn *infra.BTIntNode) *infra.BTIntNode {

	s := stack.New()

	for tn != nil {
		s.Push(tn)
		tn = tn.Left
	}

	var newHead *infra.BTIntNode
	for s.Len() != 0 {
		node := s.Pop().(*infra.BTIntNode)

		if newHead == nil {
			newHead = node
		}
		if s.Len() != 0 {
			node.Left = s.Peek().(*infra.BTIntNode).Right
			node.Right = s.Peek().(*infra.BTIntNode)
		} else {
			node.Left = nil
			node.Right = nil
		}

	}

	return newHead
}
