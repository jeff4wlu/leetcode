package leet108

import "leetcode/infra"

func SortedArrBst(in []int) *infra.BTIntNode {

	res := &infra.BTIntNode{}
	dfs(in, res)
	return res

}

func dfs(in []int, tn *infra.BTIntNode) {

	n := len(in)
	if n == 0 {
		return
	}

	//mid是索引
	mid := (n+1)/2 - 1

	left := in[:mid]
	if len(left) != 0 {
		tn.Left = &infra.BTIntNode{}
		dfs(left, tn.Left)
	}

	tn.Value = in[mid]

	right := in[mid+1:]
	if len(right) != 0 {
		tn.Right = &infra.BTIntNode{}
		dfs(right, tn.Right)
	}

}
