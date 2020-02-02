package leet174

import "leetcode/infra"

//DP
//https://blog.csdn.net/yysave/article/details/90954499
func DungeonGame(dungeon [][]int) int {
	row := len(dungeon)
	col := len(dungeon[0])

	dp := [][]int{}
	for i := 0; i < row; i++ {
		tmp := make([]int, col)
		dp = append(dp, tmp)
	}

	//初始化dp表，dp[i][j]表示为了达到右下角，至少有多少血量能够在行走过程中至少保持一滴血
	dp[row-1][col-1] = infra.IntMax(1, 1-dungeon[row-1][col-1])
	for i := row - 2; i >= 0; i-- {
		dp[i][col-1] = infra.IntMax(1, dp[i+1][col-1]-dungeon[i][col-1])
	}
	for i := col - 2; i >= 0; i-- {
		dp[row-1][i] = infra.IntMax(1, dp[row-1][i+1]-dungeon[row-1][i])
	}

	//要从右下角开始向上推，因为只有在右下角知道最少还剩一滴血，然后向上推，最小血量
	for i := row - 2; i >= 0; i-- {
		for j := col - 2; j >= 0; j-- {
			dp[i][j] = infra.IntMax(1, infra.IntMin(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
		}
	}

	return dp[0][0]
}
