package leet99

import (
	"leetcode/infra"
	"sort"
)

func RecoverBST(tn *infra.BTIntNode) {
	path := []int{}
	tnodes := []*infra.BTIntNode{}
	inorder(tn, &tnodes, &path)
	sort.Ints(path)
	for i := 0; i < len(tnodes); i++ {
		tnodes[i].Value = path[i]
	}
}

func inorder(tn *infra.BTIntNode, tnodes *[]*infra.BTIntNode, path *[]int) {
	if tn == nil {
		return
	}

	inorder(tn.Left, tnodes, path)
	*path = append(*path, tn.Value)
	*tnodes = append(*tnodes, tn)
	inorder(tn.Right, tnodes, path)

}
