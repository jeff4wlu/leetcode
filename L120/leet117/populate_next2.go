package leet117

type treeLinkNode struct {
	val   int
	left  *treeLinkNode
	right *treeLinkNode
	next  *treeLinkNode
}

func PopulateNext2(tn *treeLinkNode) {
	helper(tn)
}

func helper(tn *treeLinkNode) {

	if tn == nil || tn.left == nil && tn.right == nil {
		return
	}

	parent := tn
	var nextNode *treeLinkNode
	for parent.next != nil {
		if parent.next.left != nil {
			nextNode = parent.next.left
			break
		}
		if parent.next.right != nil {
			nextNode = parent.next.right
			break
		}
		parent = parent.next
	}

	if tn.left != nil && tn.right != nil {
		tn.left.next = tn.right
		tn.right.next = nextNode
	} else if tn.right != nil {
		tn.right.next = nextNode
	} else if tn.left != nil {
		tn.left.next = nextNode
	}

	helper(tn.left)
	helper(tn.right)

}
