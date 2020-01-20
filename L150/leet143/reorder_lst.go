package leet143

import "leetcode/infra"

func ReorderLst(tn *infra.ListNode) *infra.ListNode {

	//找出中点
	dummy := tn
	var cnt int
	for dummy != nil {
		cnt++
		dummy = dummy.Next
	}

	mid := (cnt+1)/2 - 1
	dummy = tn
	for mid > 0 {
		mid--
		dummy = dummy.Next
	}

	//翻转第二链表，使用三个临时变量
	var pre, post, cur *infra.ListNode
	cur = dummy.Next
	dummy.Next = nil

	for cur != nil {
		post = cur.Next
		cur.Next = pre
		pre = cur
		cur = post
		if post != nil {
			post = post.Next
		}

	}

	//串起来
	second := pre
	dummy = tn
	for dummy != nil && second != nil {
		tmp := dummy.Next
		tmp2 := second.Next
		dummy.Next = second
		second.Next = tmp
		dummy = tmp
		second = tmp2

	}

	return tn

}
