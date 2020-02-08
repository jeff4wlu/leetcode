package leet236

import "leetcode/infra"

//dfs遍历找出两个结点，并保存过程结点，最后找出最大高度的共同结点。
func LCABt(bt *infra.BTIntNode, p, q int) int {

	if bt == nil {
		return -1
	}
	path := []int{}
	res := [][]int{}
	dfs(bt, p, q, path, &res)

	n := infra.IntMin(len(res[0]), len(res[1]))
	var i int
	for ; i < n; i++ {
		if res[0][i] != res[1][i] {
			break
		}
	}
	return res[0][i-1]
}

func dfs(bt *infra.BTIntNode, p, q int, path []int, res *[][]int) {
	if bt == nil {
		return
	}

	if bt.Value == p || bt.Value == q {
		path = append(path, bt.Value)
		(*res) = append((*res), path)
	}

	//由于每一次递归都使用了append，所以每一层path都是一份copy。不会影响其他dfs的结果。
	path = append(path, bt.Value)
	dfs(bt.Left, p, q, path, res)
	dfs(bt.Right, p, q, path, res)

}
