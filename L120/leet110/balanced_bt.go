package leet110

import "leetcode/L120/leet104"

import "leetcode/infra"

func BalancedBT(tn *infra.BTIntNode) bool {
	a := leet104.MaxDepthBT(tn.Left)
	b := leet104.MaxDepthBT(tn.Right)

	if a-b > 1 || b-a > 1 {
		return false
	}

	return true
}
