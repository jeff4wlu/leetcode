package leet173

import "leetcode/infra"
import "github.com/golang-collections/collections/stack"

//题意是针对二叉搜索树，他特点是左子树比结点小，右子树比结点大。每次next要找出当前最小值。
//所以我们使用栈来迭代，栈顶保存最小值。二叉搜索树的最左是最小，最右是最大。
//压栈顺序是前缀遍历，先压结点再压左树

type bstIterator struct {
	bst *infra.BTIntNode
	s   *stack.Stack
}

func New(tn *infra.BTIntNode) *bstIterator {

	res := new(bstIterator)
	res.bst = tn
	res.s = stack.New()
	res.fill(tn)
	return res

}

func (bst *bstIterator) next() int {
	if bst.hasNext() {
		res := bst.s.Pop().(*infra.BTIntNode)
		if res.Right != nil {
			bst.fill(res.Right)
		}

		return res.Value
	}
	return -1
}

func (bst *bstIterator) fill(root *infra.BTIntNode) {
	for root != nil {
		bst.s.Push(root)
		root = root.Left
	}
}

func (bst *bstIterator) hasNext() bool {
	return bst.s.Len() != 0
}
