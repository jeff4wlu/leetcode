package base

type ListNode struct {
	Value int
	Next  *ListNode
}

func MakeListNode(nums []int) *ListNode {

	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Value: nums[0],
	}

	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Value: nums[i]}
		temp = temp.Next
	}

	return res
}
