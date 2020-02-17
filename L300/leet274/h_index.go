package leet274

import "leetcode/infra"

//了解题意后就可以利用排序来解决
func HIndex(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	cpy := infra.BubbleSort(nums, false)

	max := infra.INT32_MIN
	for i := 0; i < n; i++ {
		if n-i >= cpy[i] {
			max = infra.IntMax(max, cpy[i])
		}
	}

	return max
}
