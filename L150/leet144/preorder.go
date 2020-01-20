package leet144

import "leetcode/infra"

func PreOrder(tn *infra.BTIntNode) []int {

	path := []int{}
	dfs(tn, &path)
	return path
}

func dfs(tn *infra.BTIntNode, path *[]int) {
	if tn == nil {
		return
	}

	*path = append(*path, tn.Value)
	dfs(tn.Left, path)
	dfs(tn.Right, path)
}
