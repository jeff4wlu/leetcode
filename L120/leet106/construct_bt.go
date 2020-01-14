package leet106

import "leetcode/infra"

func ConstructBT(pos, in []int) *infra.BTIntNode {
	res := infra.BTIntNode{}
	dfs(pos, in, &res)
	return &res
}

func dfs(pos, in []int, tn *infra.BTIntNode) {

	n := len(pos)
	if n == 0 {
		return
	}

	tn.Value = pos[n-1]

	i := n - 1
	for ; i >= 0; i-- {
		if in[i] == tn.Value {
			break
		}
	}

	leftIn := in[:i]
	rightIn := in[i+1:]

	l := len(leftIn)
	r := len(rightIn)

	if l != 0 {
		leftpos := pos[:0+l]
		tn.Left = &infra.BTIntNode{}
		dfs(leftpos, leftIn, tn.Left)
	}

	if r != 0 {
		rightpos := pos[0+l : n-1]
		tn.Right = &infra.BTIntNode{}
		dfs(rightpos, rightIn, tn.Right)
	}
}
