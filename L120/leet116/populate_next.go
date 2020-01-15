package leet116

type treeLinkNode struct {
	val   int
	left  *treeLinkNode
	right *treeLinkNode
	next  *treeLinkNode
}

//如果有空间复杂度为O(1)要求时，只能用这种方式
func PopulateNext(tn *treeLinkNode) {

	if tn == nil {
		return
	}

	start := tn

	for start.left != nil {
		cur := start
		for cur != nil {
			cur.left.next = cur.right
			if cur.next != nil {
				cur.right.next = cur.next.left

			}
			cur = cur.next
		}
		start = start.left
	}

}
