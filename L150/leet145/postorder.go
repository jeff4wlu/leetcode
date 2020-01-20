package leet145

import "leetcode/infra"

func PostOrder(tn *infra.BTIntNode) []int {
	path := []int{}
	dfs(tn, &path)
	return path
}

func dfs(tn *infra.BTIntNode, path *[]int) {
	if tn == nil {
		return
	}

	dfs(tn.Left, path)
	dfs(tn.Right, path)
	*path = append(*path, tn.Value)
}
