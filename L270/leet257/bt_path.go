package leet257

import (
	"leetcode/infra"
	"strconv"
)

func BTPath(root *infra.BTIntNode) []string {
	if root == nil {
		return []string{}
	}

	res := [][]int{}
	path := []int{}
	dfs(root, path, &res)
	return toStrings(res)
}

func dfs(root *infra.BTIntNode, path []int, res *[][]int) {

	if root == nil {
		return
	}

	path = append(path, root.Value)

	if root.Left == nil && root.Right == nil {
		(*res) = append((*res), path)
		return
	}

	dfs(root.Left, path, res)
	dfs(root.Right, path, res)
}

func toStrings(nums [][]int) []string {

	row := len(nums)
	res := []string{}
	for i := 0; i < row; i++ {
		var tmp string
		for j := 0; j < len(nums[i]); j++ {
			tmp += strconv.Itoa(nums[i][j])
			tmp += "->"
		}
		tmp = tmp[:len(tmp)-2]
		res = append(res, tmp)
	}
	return res
}
