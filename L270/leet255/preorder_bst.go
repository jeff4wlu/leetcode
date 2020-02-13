package leet255

import "leetcode/infra"

func PreorderBST(root *infra.BTIntNode, seq []int) bool {

	n := len(seq)
	res := []int{}
	preorderDfs(root, &res)
	for i := 0; i < n; i++ {
		if seq[i] != res[i] {
			return false
		}

	}

	return true

}

func preorderDfs(root *infra.BTIntNode, path *[]int) {
	if root == nil {
		return
	}

	(*path) = append((*path), root.Value)
	preorderDfs(root.Left, path)
	preorderDfs(root.Right, path)
}
