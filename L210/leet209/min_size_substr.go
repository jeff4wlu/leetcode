package leet209

import "leetcode/infra"

func MinSizeSubstr(nums []int, k int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	var left, right, sum int
	min := infra.INT32_MAX

	for right < n {
		sum += nums[right]
		for sum >= k {
			min = infra.IntMin(min, right-left+1)
			if left == right {
				break
			}
			sum -= nums[left]
			left++
		}
		right++
	}

	if min != infra.INT32_MAX {
		return min
	}
	return 0
}
