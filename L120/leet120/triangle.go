package leet120

import "leetcode/infra"

//递归，即穷举所有可能，计算出最小值。
//穷举即遍历所有路径，过程中会返回当前路径的和
func Triangle(in [][]int) int {
	return dfs(in, 0, 0, in[0][0], infra.INT32_MAX)
}

//x,y是index
func dfs(in [][]int, x, y, cur, min int) int {

	row := len(in)

	if x == row-1 {
		if cur < min {
			return cur
		}
		return min
	}

	//路径上每个节点的下一步都只有两种可能
	if x+1 >= 0 && x+1 < row {
		a := dfs(in, x+1, y, cur+in[x+1][y], min)
		b := dfs(in, x+1, y+1, cur+in[x+1][y+1], min)
		return infra.IntMin(a, b)
	}

	return min

}
