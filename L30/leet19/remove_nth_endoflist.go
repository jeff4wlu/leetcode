package leet19

import "leetcode/infra"

//注意这个链表的头个node是有效数据，而不是head
//nth假设永远有效
func RemoveNthEndofList(lst *infra.ListNode, nth int) *infra.ListNode {

	if lst == nil {
		return nil
	}

	head := lst
	cur := lst
	pre := lst
	for i := 0; i < nth; i++ {
		cur = cur.Next
	}

	if cur == nil {
		return lst.Next
	}

	for cur.Next != nil {
		cur = cur.Next
		pre = pre.Next
	}

	pre.Next = pre.Next.Next

	return head

}
