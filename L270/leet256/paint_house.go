package leet256

import "leetcode/infra"

func PaintHouse(costs [][]int) int {

	row := len(costs)

	dp := [][]int{}
	for i := 0; i < row; i++ {
		tmp := make([]int, 3)
		dp = append(dp, tmp)
	}

	for i := 0; i < 3; i++ {
		dp[0][i] = costs[0][i]
	}

	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][1] + costs[i][0]
		dp[i][2] = dp[i-1][1] + costs[i][2]
		dp[i][1] = infra.IntMin(dp[i-1][0], dp[i-1][2]) + costs[i][1]
	}

	min := infra.INT32_MAX
	for i := 0; i < 3; i++ {
		min = infra.IntMin(min, dp[row-1][i])
	}

	return min
}
