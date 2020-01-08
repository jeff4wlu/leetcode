package leet86

import "leetcode/infra"

func PartitionList(lst *infra.ListNode, x int) *infra.ListNode {

	l, r, head := infra.ListNode{}, infra.ListNode{}, infra.ListNode{}
	head.Next = lst
	idx, lIdx, rIdx := &head, &l, &r
	for idx.Next != nil {
		if idx.Next.Value < x {
			lIdx.Next = idx.Next
			lIdx = lIdx.Next
		} else {
			rIdx.Next = idx.Next
			rIdx = rIdx.Next
		}
		idx = idx.Next
	}
	rIdx.Next = nil //由于在原有链条上操作，可能会残余node在最后。

	lIdx.Next = r.Next

	return l.Next

}
