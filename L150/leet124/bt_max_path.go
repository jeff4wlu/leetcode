package leet124

import "leetcode/infra"

//dfs
func BTMaxPath(tn *infra.BTIntNode) int {
	return helper(tn)
}

func helper(tn *infra.BTIntNode) int {

	if tn == nil {
		return 0
	}

	left := infra.IntMax(0, helper(tn.Left))
	right := infra.IntMax(0, helper(tn.Right))

	whole := left + right + tn.Value

	if left == 0 && right == 0 {
		return whole
	} else if left != 0 && right != 0 {
		return infra.IntMax(infra.IntMax(left, right), whole)
	} else if left != 0 {
		return infra.IntMax(whole, left)
	}

	return infra.IntMax(whole, right)

}
