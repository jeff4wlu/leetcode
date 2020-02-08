package leet237

import "leetcode/infra"

//这道题的题意要理解清楚，它只提供要被删除结点，而不是链表的起点
//所以给定的node肯定不是最后一个结点
func DelNode(node *infra.ListNode) {

	if node == nil {
		return
	}

	for {
		node.Value = node.Next.Value
		if node.Next.Next == nil {
			node.Next = nil
			return
		}
		node = node.Next
	}

}
