package leet95

import "leetcode/infra"

func UniqueBST(n int) []*infra.BTIntNode {

	return traversal(1, n)

}

func traversal(start, end int) []*infra.BTIntNode {

	res := []*infra.BTIntNode{}
	if start > end {
		return append(res, nil)
	}

	for i := start; i <= end; i++ {

		left := traversal(start, i-1)
		right := traversal(i+1, end)
		for _, lNode := range left {
			for _, rNode := range right {
				tmp := &infra.BTIntNode{}
				tmp.Value = i
				tmp.Left = lNode
				tmp.Right = rNode
				res = append(res, tmp)
			}
		}

	}
	return res

}
