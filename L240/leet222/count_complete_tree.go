package leet222

import "leetcode/infra"

import "math"

//利用完全二叉树的特点计算
func CountCompleteTree(root *infra.BTIntNode) int {
	return dfs(root)
}

func dfs(root *infra.BTIntNode) int {
	if root == nil {
		return 0
	}

	//求最左和最右的高度
	var left, right int
	forword := root
	for forword != nil {
		left++
		forword = forword.Left
	}
	forword = root
	for forword != nil {
		right++
		forword = forword.Right
	}
	if left == right { //完全二叉树
		return int(math.Pow(2, float64(left)) - 1)
	}

	return dfs(root.Left) + dfs(root.Right) + 1

}
