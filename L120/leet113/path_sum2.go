package leet113

import "leetcode/infra"

import "fmt"

func PathSum2(tn *infra.BTIntNode, sum int) [][]int {
	res := [][]int{}
	path := []int{}
	dfs(tn, 0, sum, path, &res)
	return res
}

func dfs(tn *infra.BTIntNode, cur, sum int, path []int, res *[][]int) {

	if tn == nil {
		return
	}

	added := cur + tn.Value
	path = append(path, tn.Value)

	if tn.Left == nil && tn.Right == nil {
		fmt.Printf(" %d", tn.Value)
		if added == sum {
			tmp := make([]int, len(path))
			copy(tmp, path)
			*res = append(*res, tmp)
		}
		return
	}

	//以下代码行为是剪枝，加快递归速度
	if added >= sum {
		return
	}

	dfs(tn.Left, added, sum, path, res)
	dfs(tn.Right, added, sum, path, res)

}
