package leet148

import "leetcode/infra"

//使用归并排序，适合链表
func SortLst(ln *infra.ListNode) *infra.ListNode {
	return dfs(ln)
}

func dfs(ln *infra.ListNode) *infra.ListNode {

	if ln == nil {
		return nil
	}

	//求中点
	fast, slow, prePos := ln, ln, ln

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prePos = slow
		slow = slow.Next
	}

	if slow == ln {
		return slow
	}
	prePos.Next = nil //断开链表

	left := dfs(ln)
	right := dfs(slow)

	if left == nil && right == nil {
		return nil
	}

	if left == nil {
		return right
	} else if right == nil {
		return left
	}

	dummyHead := new(infra.ListNode)
	tmp := dummyHead
	for {
		if left.Value <= right.Value {
			tmp.Next = left
			left = left.Next
		} else {
			tmp.Next = right
			right = right.Next
		}
		tmp = tmp.Next

		if left == nil && right == nil {
			return dummyHead.Next
		}

		if left == nil {
			tmp.Next = right
			break
		} else if right == nil {
			tmp.Next = left
			break
		}
	}

	return dummyHead.Next

}
