package leet112

import "leetcode/infra"

func PathSum(tn *infra.BTIntNode, sum int) bool {
	if tn == nil && sum == 0 {
		return true
	}
	return dfs(tn, 0, sum)
}

func dfs(tn *infra.BTIntNode, cur, sum int) bool {

	if tn == nil {
		return false
	}

	added := cur + tn.Value

	if tn.Left == nil && tn.Right == nil {
		if added == sum {
			return true
		}
		return false
	}

	if added >= sum {
		return false
	}

	return dfs(tn.Left, added, sum) || dfs(tn.Right, added, sum)

}
