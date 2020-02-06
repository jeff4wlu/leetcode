package leet220

import (
	"leetcode/infra"
	"math"
)

func ContainDuplicate3(nums []int, t, k int) bool {

	n := len(nums)
	if n <= 1 || k < 1 {
		return false
	}

	for i := 0; i < n; i++ {
		end := infra.IntMin(i+k, n-1)
		for j := i + 1; j <= end; j++ {
			gap := math.Abs(float64(nums[i] - nums[j]))
			if int(gap) <= t {
				return true
			}
		}
	}

	return false
}
