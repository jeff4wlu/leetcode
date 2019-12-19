package leet23

import (
	"leetcode/L30/leet21"
	"leetcode/infra"
)

//这种算法在协程和多CPU core下性能才会体现，并行计算
func MergeKSortedList(in []*infra.ListNode) *infra.ListNode {

	n := len(in)

	for n > 1 {
		k := (n + 1) / 2 //打对折
		for i := 0; i < n/2; i++ {
			in[i] = leet21.MergeTwoSortedList(in[i], in[i+k]) //支持其中一个是空数组
		}
		n = k
	}

	return in[0]

}
