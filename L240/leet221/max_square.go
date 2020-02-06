package leet221

import "leetcode/infra"

//DP
func MaxSquare(m [][]int) int {

	row := len(m)
	col := len(m[0])

	dp := [][]int{}
	for i := 0; i < row; i++ {
		c := make([]int, col)
		dp = append(dp, c)
	}

	for i := 0; i < row; i++ {
		dp[i][0] = m[i][0]
	}
	for i := 0; i < col; i++ {
		dp[0][i] = m[0][i]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			dp[i][j] = infra.IntMin(infra.IntMin(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + m[i][j]
		}
	}

	return dp[row-1][col-1] * dp[row-1][col-1]
}
