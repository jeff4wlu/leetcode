package leet206

import "leetcode/infra"

//迭代
func ReverseLinkedLst(lst *infra.ListNode) *infra.ListNode {

	if lst == nil {
		return nil
	}
	if lst.Next == nil {
		return lst
	}

	dummyHead := new(infra.ListNode)
	dummyHead.Next = lst
	//rest要处理的串
	//pre，用来获取要处理串的最后一个Node
	//copy，用来步进已完成的串
	rest, pre, copy := lst, lst, dummyHead

	for rest != nil {

		if pre.Next == nil {
			copy.Next = pre
			break
		}
		//取要处理lst的最后一结点
		for pre.Next.Next != nil {
			pre = pre.Next
		}
		copy.Next = pre.Next
		pre.Next = nil
		pre = rest
		copy = copy.Next
	}

	return dummyHead.Next
}
