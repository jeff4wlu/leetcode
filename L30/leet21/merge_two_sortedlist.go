package leet21

import "leetcode/infra"

func MergeTwoSortedList(l1, l2 *infra.ListNode) *infra.ListNode {

	head := &infra.ListNode{
		Value: -1,
		Next:  nil,
	}

	res := head

	for (l1 == nil && l2 == nil) == false {
		if l1 != nil && l2 != nil {
			if l1.Value < l2.Value {
				head.Next = l1
				l1 = l1.Next

			} else {
				head.Next = l2
				l2 = l2.Next
			}
		} else if l1 == nil {
			head.Next = l2
			l2 = l2.Next
		} else {
			head.Next = l1
			l1 = l1.Next
		}

		head = head.Next
	}

	return res.Next

}
