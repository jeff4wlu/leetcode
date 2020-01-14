package leet105

import "leetcode/infra"

func ConstructBT(pre, in []int) *infra.BTIntNode {
	res := infra.BTIntNode{}
	dfs(pre, in, &res)
	return &res
}

func dfs(pre, in []int, tn *infra.BTIntNode) {

	if len(pre) == 0 {
		return
	}

	tn.Value = pre[0]

	i := 0
	for ; i < len(pre); i++ {
		if in[i] == tn.Value {
			break
		}
	}

	leftIn := in[:i]
	rightIn := in[i+1:]

	l := len(leftIn)
	r := len(rightIn)

	if l != 0 {
		leftPre := pre[1 : 1+l]
		tn.Left = &infra.BTIntNode{}
		dfs(leftPre, leftIn, tn.Left)
	}

	if r != 0 {
		rightPre := pre[1+l:]
		tn.Right = &infra.BTIntNode{}
		dfs(rightPre, rightIn, tn.Right)
	}
}
