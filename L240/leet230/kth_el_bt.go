package leet230

import "leetcode/infra"

//题意默认k是有效范围内的。使用搜索二叉树排序
func KthElBT(bst *infra.BTIntNode, k int) int {

	if bst == nil {
		return -1
	}

	stack := infra.CreateStack()
	node := bst
	arr := []int{}

	for node != nil || !stack.IsEmpty() {

		//遍历到最小值
		for node != nil {
			stack.Push(node)
			node = node.Left //最小到大的排序，使用中序
		}
		node = stack.Pop().(*infra.BTIntNode)
		arr = append(arr, node.Value)
		node = node.Right

	}

	return arr[k-1]
}
