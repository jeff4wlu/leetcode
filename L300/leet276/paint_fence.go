package leet276

//多少种可能性用DP，排列组合的思想，“要么”用加法
func PaintFence(n, k int) int {

	if n <= 0 || k <= 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	dp := make([]int, n+1)
	dp[1] = k
	dp[2] = k * k

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1]*(k-1) + dp[i-2]*(k-1)
	}

	return dp[n]
}
