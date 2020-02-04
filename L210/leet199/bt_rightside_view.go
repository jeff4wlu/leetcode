package leet199

import "leetcode/infra"

//用到的知识点，二叉树层序，并知道是哪一层
func BTRightView(root *infra.BTIntNode) []int {

	res := []int{}
	if root == nil {
		return res
	}

	q := infra.NewQueue()
	q.Enqueue(root)
	var nilV *infra.BTIntNode
	q.Enqueue(nilV)

	tmp := []int{}
	for q.IsEmpty() == false {
		node := q.Dequeue().(*infra.BTIntNode)
		if node == nil { //nil表示当前队列内全是下一层的结点，所以加nil作为下一层的分隔。
			if q.IsEmpty() == false {
				q.Enqueue(nilV)
			}
			if len(tmp) != 0 {
				res = append(res, tmp[0])
				tmp = []int{}
			}
			continue
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		tmp = append(tmp, node.Value)
	}

	return res
}
