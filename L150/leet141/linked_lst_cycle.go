package leet141

type lstNode struct {
	val  int
	next *lstNode
}

func LinkedLstCycle(in []int, pos int) bool {

	//构造链表
	var next, head, p *lstNode
	tmp := new(lstNode)
	tail := tmp
	for i := len(in) - 1; i >= 0; i-- {
		tmp.val = in[i]
		tmp.next = next
		next = tmp
		if i == 0 {
			head = tmp
		}
		if i == pos {
			p = tmp
		}
		tmp = new(lstNode)
	}
	tail.next = p

	//检查是否有循环
	slow, fast := head, head

	for slow != nil && fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}

	}
	return false
}
