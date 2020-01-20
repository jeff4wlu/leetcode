package leet141

import "leetcode/infra"

func LinkedLstCycle(in []int, pos int) bool {

	//构造链表
	var post, head, p *infra.ListNode
	tmp := new(infra.ListNode)
	tail := tmp
	for i := len(in) - 1; i >= 0; i-- {
		tmp.Value = in[i]
		tmp.Next = post
		post = tmp
		if i == 0 {
			head = tmp
		}
		if i == pos {
			p = tmp
		}
		tmp = new(infra.ListNode)
	}
	tail.Next = p

	//检查是否有循环
	slow, fast := head, head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}

	}
	return false
}
