package leet45

import "leetcode/infra"

//本算法使用的是BFS深度优先遍历
func JumpGameII(nums []int) int {

	n := len(nums)

	var curPos, level, i int

	for curPos < n-1 {

		level++ //BFS的level
		prePos := curPos
		for ; i <= prePos; i++ {
			//求出当前为止能够跳到的最远范围，而不是当前为止的最远范围。当前最远是包含了前面几跳的累积最远效果
			curPos = infra.IntMax(curPos, i+nums[i])
		}

	}

	return level

}
