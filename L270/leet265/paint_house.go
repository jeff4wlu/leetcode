package leet265

import "leetcode/infra"

//记录最小值和次小值
func PaintHouse(costs [][]int) int {

	row := len(costs)
	col := len(costs[0])

	dp := [][]int{}
	for i := 0; i < row; i++ {
		tmp := make([]int, col)
		dp = append(dp, tmp)
	}

	min1, min2 := infra.INT32_MAX, infra.INT32_MAX
	for i := 0; i < col; i++ {
		dp[0][i] = costs[0][i]
		if costs[0][i] < min1 {
			min2 = min1
			min1 = costs[0][i]
		} else if costs[0][i] < min2 {
			min2 = costs[0][i]
		}
	}

	for i := 1; i < row; i++ {
		tmp1, tmp2 := infra.INT32_MAX, infra.INT32_MAX
		for j := 0; j < col; j++ {
			if min1 == dp[i-1][j] {
				dp[i][j] = min2 + costs[i][j]
			} else {
				dp[i][j] = min1 + costs[i][j]
			}
			if dp[i][j] < tmp1 {
				tmp2 = tmp1
				tmp1 = dp[i][j]
			} else if dp[i][j] < tmp2 {
				tmp2 = dp[i][j]
			}
		}
		min1, min2 = tmp1, tmp2
	}
	return min1
}

/*
func getTwoMin(nums []int) (min1, min2 int) {

	min1, min2 = infra.INT32_MAX, infra.INT32_MAX
	n := len(nums)
	if n < 2 {
		return
	}

	for i := 0; i < n; i++ {
		if nums[i] < min1 {
			min2 = min1
			min1 = nums[i]
		} else if nums[i] < min2 {
			min2 = nums[i]
		}
	}
	return
}*/
