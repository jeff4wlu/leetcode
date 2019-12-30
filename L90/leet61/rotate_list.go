package leet61

import "leetcode/infra"

func RotateList(lst *infra.ListNode, k int) *infra.ListNode {

	var idx, end, res *infra.ListNode
	size := 1

	copy := lst
	for {
		if lst.Next == nil {
			end = lst
			break
		}
		size++
		lst = lst.Next
	}

	step := size - k%size

	//idx是链表头的前一个node
	idx = copy
	for i := 0; i < step-1; i++ {
		idx = idx.Next
	}

	//求出链表头
	res = idx.Next
	idx.Next = nil
	end.Next = copy

	return res

}
