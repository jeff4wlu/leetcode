package leet114

import "leetcode/infra"

func FlatternBT2List(tn *infra.BTIntNode) {
	helper(tn)
}

func helper(tn *infra.BTIntNode) {

	if tn == nil {
		return
	}

	var leftParent *infra.BTIntNode
	left := tn //找出最左节点的父节点
	for left.Left != nil {
		leftParent = left
		left = left.Left
	}

	if leftParent != nil {
		tmp := left
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = leftParent.Right
		leftParent.Right = left
		leftParent.Left = nil
		helper(tn)
	}

}
