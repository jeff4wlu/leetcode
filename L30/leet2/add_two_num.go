package leet2

import "leetcode/infra"

func AddTwoNum(l1, l2 *infra.ListNode) (re *infra.ListNode) {

	dummy := infra.ListNode{-1, nil}
	cur := &dummy
	var carry int

	for l1 != nil || l2 != nil {
		var v1 int
		var v2 int
		if l1 != nil {
			v1 = l1.Value
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Value
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		if sum >= 10 {
			carry = 1
		}
		cur.Next = &infra.ListNode{sum % 10, nil}
		cur = cur.Next

	}

	if carry == 1 {
		cur.Next = &infra.ListNode{1, nil}
	}

	return dummy.Next

}
