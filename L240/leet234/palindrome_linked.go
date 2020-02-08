package leet234

import "leetcode/infra"

//链表必须两个元素或以上
func PalindromeLinked(lst *infra.ListNode) bool {

	if lst == nil {
		return false
	}

	//求链表中点，以中点为开始切断链表为两条
	var pre *infra.ListNode
	fast, slow := lst, lst
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil

	//反转第二段链表。需要使用node,newHead两个辅助指针。最后newhead就是反转的链表
	var node, newHead *infra.ListNode
	head := slow
	for head != nil {
		node = head
		head = head.Next
		node.Next = newHead
		newHead = node
	}

	//比较前后两端链表元素是否一样
	for lst != nil {
		if lst.Value != newHead.Value {
			return false
		}
		lst = lst.Next
		newHead = newHead.Next
	}

	return true
}
