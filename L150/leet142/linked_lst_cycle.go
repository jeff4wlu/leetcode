package leet142

type lstNode struct {
	val  int
	next *lstNode
}

func LinkedLstCycle(in []int, pos int) (start, res *lstNode) {

	//构造链表
	var next, p *lstNode
	tmp := new(lstNode)
	tail := tmp
	head := new(lstNode)
	for i := len(in) - 1; i >= 0; i-- {
		tmp.val = in[i]
		tmp.next = next
		next = tmp
		if i == 0 {
			head.next = tmp
		}
		if i == pos {
			p = tmp
		}
		tmp = new(lstNode)
	}
	tail.next = p

	//检查是否有循环
	slow, fast := head, head
	var met *lstNode
	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			met = slow
			break
		}

	}

	//这里不需要两倍的速度，而且需要虚拟头指针
	if met != nil {
		slow, fast = head, met
		for slow != nil {
			slow = slow.next
			fast = fast.next
			if slow == fast {
				met = slow
				break
			}

		}
	}

	return p, met
}
