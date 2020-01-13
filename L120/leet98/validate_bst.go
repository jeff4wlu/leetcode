package leet98

import "leetcode/infra"

//使用自身性质，且使用递归
func ValidateBST(tn *infra.BTIntNode) bool {
	return helper(tn, infra.INT32_MIN, infra.INT32_MAX)
}

func helper(tn *infra.BTIntNode, min, max int) bool {

	if tn == nil {
		return true
	}

	if tn.Value < min || tn.Value > max {
		return false
	}

	return helper(tn.Left, min, tn.Value) && helper(tn.Right, tn.Value, max)

}

func ValidateBST1(tn *infra.BTIntNode) bool {
	path := []int{}
	inorder(tn, &path)
	for i := 0; i < len(path)-1; i++ {
		if path[i] > path[i+1] {
			return false
		}
	}
	return true
}

func inorder(tn *infra.BTIntNode, path *[]int) {
	if tn == nil {
		return
	}

	inorder(tn.Left, path)
	*path = append(*path, tn.Value)
	inorder(tn.Right, path)

}
