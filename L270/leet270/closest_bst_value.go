package leet270

import "leetcode/infra"

var min float32 = infra.INT32_MAX
var value int = -1

func ClosestValue(root *infra.BTIntNode, target float32) int {
	dfs(root, target)
	return value
}

//bst是中序遍历
func dfs(root *infra.BTIntNode, target float32) {

	if root == nil {
		return
	}

	diff := target - float32(root.Value)
	if diff < 0 {
		diff = -diff
		if float32(min) > diff {
			min = diff
			value = root.Value
		}
		dfs(root.Left, target)
	} else {
		if float32(min) > diff {
			min = diff
			value = root.Value
		}
		dfs(root.Right, target)
	}
}
