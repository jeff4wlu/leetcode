package leet235

import "leetcode/infra"

//需要利用二叉搜索树的特点来优化二叉树的遍历。
//为了简便，保证p比q小。且使用-1代表失败
//最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，
//最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
func LCABt(bt *infra.BTIntNode, p, q int) int {

	if q < p {
		p, q = q, p
	}
	return dfs(bt, p, q)
}

func dfs(bt *infra.BTIntNode, p, q int) int {
	if bt == nil {
		return -1
	}

	if bt.Value > q {
		return dfs(bt.Left, p, q)
	} else if bt.Value < p {
		return dfs(bt.Right, p, q)
	}

	return bt.Value

}
