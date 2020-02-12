package leet250

import "leetcode/infra"

var count int

//后续，从下到上
func CountSubtrees(root *infra.BTIntNode) int {
	dfs(root)
	return count
}

func dfs(root *infra.BTIntNode) bool {
	if root == nil {
		return true
	} else if root.Left == nil && root.Right == nil {
		count++
		return true
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	if root.Left != nil && root.Right != nil {
		if left && right && root.Value == root.Left.Value && root.Value == root.Right.Value {
			count++
			return true
		}
		return false
	} else if root.Left != nil {
		if left && root.Left.Value == root.Value {
			count++
			return true
		}
		return false
	} else {
		if right && root.Right.Value == root.Value {
			count++
			return true
		}
	}

	return false

}
