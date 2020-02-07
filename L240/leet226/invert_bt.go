package leet226

import "leetcode/infra"

//使用二叉树的层序遍历，利用队列来递推
func InvertBT(bt *infra.BTIntNode) *infra.BTIntNode {

	if bt == nil {
		return nil
	}

	q := infra.NewQueue()
	q.Enqueue(bt)
	for !q.IsEmpty() {
		btNode := q.Dequeue().(*infra.BTIntNode)
		if btNode != nil {
			q.Enqueue(btNode.Left)
			q.Enqueue(btNode.Right)
			btNode.Left, btNode.Right = btNode.Right, btNode.Left
		}
	}

	return bt
}

func InvertBTDFS(bt *infra.BTIntNode) *infra.BTIntNode {
	return dfs(bt)
}

//使用递归
func dfs(root *infra.BTIntNode) *infra.BTIntNode {

	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = dfs(root.Right)
	root.Right = dfs(tmp)

	return root
}
