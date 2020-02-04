package infra

type ListNode struct {
	Value int
	Next  *ListNode
}

type DListNode struct {
	Value int
	Next  *DListNode
	Pre   *DListNode
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

func CompareTwoIntList(got, want *ListNode) bool {
	if got == nil && want == nil {
		return true
	}
	if got == nil || want == nil {
		return false
	}

	for got != nil && want != nil {
		if got.Value != want.Value {
			return false
		}
		got = got.Next
		want = want.Next
		if got != nil && want == nil || got == nil && want != nil {
			return false
		}
	}
	return true
}
