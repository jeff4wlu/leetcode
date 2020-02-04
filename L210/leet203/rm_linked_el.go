package leet203

import "leetcode/infra"

func RMLinkedEl(lst *infra.ListNode, k int) *infra.ListNode {

	dummyHead := new(infra.ListNode)
	dummyHead.Next = lst
	pre, forward := dummyHead, dummyHead.Next

	for forward != nil {
		if forward.Value == k {
			pre.Next = forward.Next
			forward = forward.Next
			continue
		}
		forward = forward.Next
		pre = pre.Next
	}

	return dummyHead.Next
}
