package add_two_num

import "leetcode/base"

func AddTwoNum(l1, l2 *base.ListNode) (re *base.ListNode) {

	dummy := base.ListNode{-1, nil}
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
		cur.Next = &base.ListNode{sum % 10, nil}
		cur = cur.Next

	}

	if carry == 1 {
		cur.Next = &base.ListNode{1, nil}
	}

	return dummy.Next

}
