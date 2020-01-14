package leet101

import "leetcode/infra"

//dfs
func SymetricTree(tn *infra.BTIntNode) bool {
	return dfs(tn.Left, tn.Right)
}

func dfs(a, b *infra.BTIntNode) bool {

	if a == nil && b == nil {
		return true
	} else if a == nil && b != nil || a != nil && b == nil || a.Value != b.Value {
		return false
	}

	if dfs(a.Left, b.Right) && dfs(a.Right, b.Left) {
		return true
	}
	return false
}
