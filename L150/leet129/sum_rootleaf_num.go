package leet129

import (
	"leetcode/infra"
	"math"
)

func SumRootLeafNum(tn *infra.BTIntNode) int {

	path := []int{}
	res := [][]int{}
	dfs(tn, path, &res)

	n := len(res)
	if n == 0 {
		return 0
	}

	var sum int
	for _, v := range res {
		l := len(v)
		for _, t := range v {
			sum += t * int(math.Pow10(l-1))
			l--
		}
	}

	return sum
}

func dfs(tn *infra.BTIntNode, path []int, res *[][]int) {

	if tn == nil {
		return
	}

	if tn.Left == nil && tn.Right == nil {
		path = append(path, tn.Value)
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	path = append(path, tn.Value)
	dfs(tn.Left, path, res)
	dfs(tn.Right, path, res)

}
