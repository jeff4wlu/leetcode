package leet25

import "leetcode/infra"

var stack infra.Stack = infra.CreateStack()

//pre是翻转链表的前一个，即dummy头；n是翻转链表的后一个元素；cur是翻转链表最后一个元素
//本算法使用栈来翻转，貌似不够简洁。
func ReverseKGroup(in *infra.ListNode, k int) *infra.ListNode {

	var dummy infra.ListNode
	dummy.Next = in
	pre := &dummy

	cur := pre
	for i := 0; i < k; i++ {
		cur = cur.Next
		if cur == nil {
			return nil
		}
	}
	n := cur.Next

	for {

		last := ReverseList(pre, k)
		last.Next = n //接上后面未翻转的链表

		for i := 0; i < k; i++ {
			pre = pre.Next
			cur = pre
			for j := 0; j < k; j++ {
				cur = cur.Next
			}
		}

		if cur == nil {
			break
		}
		n = cur.Next
	}

	return dummy.Next

}

//使用k来控制翻转多少个元素
//返回的pre是已经捋顺的链表，但还没接上后面的，所以要返回最后一个节点去接上后面的。
func ReverseList(pre *infra.ListNode, k int) *infra.ListNode {

	step := pre

	for i := 0; i < k; i++ {
		stack.Push(step.Next)
		step = step.Next
	}

	pre.Next = stack.Pop().(*infra.ListNode)
	step = pre.Next

	for stack.IsEmpty() == false {
		step.Next = stack.Pop().(*infra.ListNode)
		step = step.Next
	}

	return step

}
