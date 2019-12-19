package leet24

import "leetcode/infra"

//关键点在t这个临时变量，它要正确记录1->3->4这个顺序。因为它后面要更新pre（步进）
func SwapNodesPair(in *infra.ListNode) *infra.ListNode {

	var dummy infra.ListNode

	pre := &dummy
	dummy.Next = in

	for pre.Next != nil && pre.Next.Next != nil {
		t := pre.Next.Next
		pre.Next.Next = t.Next
		t.Next = pre.Next
		pre.Next = t
		pre = t.Next
	}

	return dummy.Next
}
