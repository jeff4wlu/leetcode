package leet100

import "leetcode/infra"

func SameTree(a, b *infra.BTIntNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil && b != nil || a != nil && b == nil || a.Value != b.Value {
		return false
	}

	return SameTree(a.Left, b.Left) && SameTree(a.Right, b.Right)
}
