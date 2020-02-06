package leet219

import "leetcode/infra"

func ContainDuplicate2(nums []int, k int) bool {

	n := len(nums)
	if n <= 1 || k < 1 {
		return false
	}

	for i := 0; i < n; i++ {
		end := infra.IntMin(i+k, n-1)
		for j := i + 1; j <= end; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
