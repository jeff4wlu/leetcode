package leet111

import "leetcode/infra"

func MinBTDepth(tn *infra.BTIntNode) int {
	return dfs(tn, 1, infra.INT32_MAX)
}

func dfs(tn *infra.BTIntNode, lvl, min int) int {

	if tn == nil {
		return 0
	}

	if tn.Left == nil && tn.Right == nil {
		return infra.IntMin(lvl, min)
	}

	left, right := infra.INT32_MAX, infra.INT32_MAX
	if tn.Left != nil {
		left = dfs(tn.Left, lvl+1, min)
	}

	if tn.Right != nil {
		right = dfs(tn.Right, lvl+1, min)
	}

	return infra.IntMin(left, right)

}
