package leet272

import "leetcode/infra"

func ClosestVal(root *infra.BTIntNode, target float32, k int) []int {

	if root == nil {
		return []int{}
	}

	res := []int{}
	q := infra.NewQueue()
	dfs(root, q, target, k)
	for !q.IsEmpty() {
		res = append(res, q.Dequeue().(int))
	}

	return res
}

func dfs(root *infra.BTIntNode, q *infra.Queue, target float32, k int) {

	if root == nil {
		return
	}
	dfs(root.Left, q, target, k)
	if q.Size() >= k {
		tmp := q.Peek().(int)
		diff1 := float32(tmp) - target
		if diff1 < 0 {
			diff1 = -diff1
		}
		diff2 := float32(root.Value) - target
		if diff2 < 0 {
			diff2 = -diff2
		}
		if diff1 > diff2 {
			q.Enqueue(root.Value)
			q.Dequeue()
		}

	} else {
		q.Enqueue(root.Value)

	}
	dfs(root.Right, q, target, k)
}
