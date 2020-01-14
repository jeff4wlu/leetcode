package leet102

import "leetcode/infra"

func BTOrderLvlTra(tn *infra.BTIntNode) [][]int {

	res := [][]int{}
	dfs(tn, 1, &res)
	return res
}

func dfs(tn *infra.BTIntNode, lvl int, res *[][]int) {

	if tn == nil {
		return
	}

	if len(*res) < lvl {
		tmp := []int{tn.Value}
		*res = append(*res, tmp)
	} else {
		(*res)[lvl-1] = append((*res)[lvl-1], tn.Value)
	}

	dfs(tn.Left, lvl+1, res)
	dfs(tn.Right, lvl+1, res)

}
