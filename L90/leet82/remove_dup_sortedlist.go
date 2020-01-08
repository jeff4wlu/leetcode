package leet82

import "leetcode/infra"

func RemoveDupSortedList(nums *infra.ListNode) *infra.ListNode {

	head := nums
	cur, next := nums, nums.Next

	for next != nil {
		for cur.Value == next.Value {
			cur.Next = next.Next
			next = cur.Next
		}
		cur = cur.Next
		next = next.Next
	}

	return head

}
