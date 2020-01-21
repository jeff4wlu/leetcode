package leet147

import "leetcode/infra"

func InsertionSortedLst(ln *infra.ListNode) *infra.ListNode {

	if ln == nil {
		return ln
	}

	dummyHead := new(infra.ListNode)
	dummyHead.Next = ln
	cur := ln

	for cur != nil {

		if cur.Next != nil && cur.Value > cur.Next.Value {

			tmp := cur.Next
			cur.Next = tmp.Next
			prePos := dummyHead
			for prePos.Next.Value < tmp.Value {
				prePos = prePos.Next
			}
			tmp.Next = prePos.Next
			prePos.Next = tmp
			continue

		}
		cur = cur.Next
	}

	return dummyHead.Next

}
