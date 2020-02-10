package leet239

import "leetcode/infra"

func SlidingWin(nums []int, k int) []int {

	res := []int{}

	n := len(nums)
	if k > n || n < 1 {
		return res
	}

	heap := infra.NewHeap2(false)

	for i := 0; i < k; i++ {
		heap.Insert(nums[i])
	}
	res = append(res, heap.GetTop())

	for i := k; i < n; i++ {
		heap.Remove(nums[i-k])
		heap.Insert(nums[i])
		res = append(res, heap.GetTop())
	}

	return res

}
