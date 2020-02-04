package leet198

import "leetcode/infra"

//DP
//求状态值前i家房屋所能获得的最大收益。
func HouseRobber(houses []int) int {

	n := len(houses)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return houses[1]
	}

	dp := make([]int, n)
	dp[0] = houses[0]
	dp[1] = infra.IntMax(dp[1], houses[1])

	for i := 2; i < n; i++ {
		dp[i] = infra.IntMax(dp[i-1], dp[i-2]+houses[i])
	}

	return dp[n-1]
}
