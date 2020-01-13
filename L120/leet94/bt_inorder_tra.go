package leet94

import "leetcode/infra"

func BTInorderTra(t *infra.BTIntNode)[]int {

	path:=[]int{}
	solve(t, &path)
	return path
}

func solve(t *infra.BTIntNode, path *[]int){
	if t == nil{
		return
	}

	solve(t.Left, path)
	*path = append(*path, t.Value)
	solve(t.Right, path)
}
