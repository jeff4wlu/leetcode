package leet42

import "leetcode/infra"

//存储水的条件是两高夹着一矮
func TrappingRainwater(nums []int) int {

	var res int

	n := len(nums)

	//扫描数组
	for i := 1; i < n-1; i++ {
		var maxLeft, maxRight int
		for j := i; j >= 0; j-- {
			maxLeft = infra.IntMax(nums[j], maxLeft)
		}
		for j := i; j < n; j++ {
			maxRight = infra.IntMax(nums[j], maxRight)
		}

		res += infra.IntMin(maxLeft, maxRight) - nums[i]
	}

	return res
}
