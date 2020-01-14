package leet104

import "leetcode/infra"

func MaxDepthBT(tn *infra.BTIntNode) int {
	return dfs(tn, 0)
}

func dfs(tn *infra.BTIntNode, max int) int {

	if tn == nil {
		return max
	}

	if tn.Left == nil && tn.Right == nil {
		return max + 1
	}

	var l, r int
	if tn.Left != nil {
		l = dfs(tn.Left, max+1)
	}
	if tn.Right != nil {
		r = dfs(tn.Right, max+1)
	}

	return infra.IntMax(l, r)
}
