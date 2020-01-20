package leet142

import "leetcode/infra"

func LinkedLstCycle(in []int, pos int) (start, res *infra.ListNode) {

	//构造链表
	var Next, p *infra.ListNode
	tmp := new(infra.ListNode)
	tail := tmp
	head := new(infra.ListNode)
	for i := len(in) - 1; i >= 0; i-- {
		tmp.Value = in[i]
		tmp.Next = Next
		Next = tmp
		if i == 0 {
			head.Next = tmp
		}
		if i == pos {
			p = tmp
		}
		tmp = new(infra.ListNode)
	}
	tail.Next = p

	//检查是否有循环
	slow, fast := head, head
	var met *infra.ListNode
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			met = slow
			break
		}

	}

	//这里不需要两倍的速度，而且需要虚拟头指针
	if met != nil {
		slow, fast = head, met
		for slow != nil {
			slow = slow.Next
			fast = fast.Next
			if slow == fast {
				met = slow
				break
			}

		}
	}

	return p, met
}
