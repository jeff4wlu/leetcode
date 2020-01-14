package leet107

import "leetcode/infra"

func BTLvlOrder(tn *infra.BTIntNode) [][]int {
	tmp := [][]int{}
	dfs(tn, 0, &tmp)

	res := [][]int{}
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}

	return res
}

func dfs(tn *infra.BTIntNode, lvl int, res *[][]int) {

	if tn == nil {
		return
	}

	if len(*res) == lvl { //路过先分配内存
		*res = append(*res, []int{})
	}

	//以下可以是先序，中序，后序都可以，不影响结果
	dfs(tn.Left, lvl+1, res)
	dfs(tn.Right, lvl+1, res)
	(*res)[lvl] = append((*res)[lvl], tn.Value)

}
