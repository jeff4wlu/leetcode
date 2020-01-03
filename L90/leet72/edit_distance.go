package leet72

import "leetcode/infra"

func EditDistance(a, b string) int {

	aLen := len(a)
	bLen := len(b)

	//外循环是第一个[],内循环是第二个[]
	dp := [][]int{}
	for i := 0; i <= aLen; i++ {
		tmp := make([]int, bLen+1)
		dp = append(dp, tmp)
	}

	dp[0][0] = 0
	for i := 1; i <= aLen; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= bLen; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= aLen; i++ {
		for j := 1; j <= bLen; j++ {
			if a[i-1:i] == b[j-1:j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}

		}
	}

	return dp[aLen][bLen]
}

func min(a, b, c int) int {
	return infra.IntMin(infra.IntMin(a, b), c)
}
