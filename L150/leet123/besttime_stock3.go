package leet123

import "leetcode/infra"

//分治
func BestTimeStock31(in []int) int {
	n := len(in)
	if n <= 1 {
		return 0
	}
	if n == 2 {
		return in[1] - in[0]
	}

	firstProfit := make([]int, n)
	secondProfit := make([]int, n)

	min := in[0]
	for i := 1; i < n; i++ {
		min = infra.IntMin(min, in[i])
		firstProfit[i] = infra.IntMax(firstProfit[i-1], in[i]-min)
	}

	max := in[n-1]
	for i := n - 2; i >= 0; i-- {
		max = infra.IntMax(max, in[i])
		secondProfit[i] = infra.IntMax(secondProfit[i+1], max-in[i])
	}

	max = 0
	for i := 0; i < n; i++ {
		max = infra.IntMax(max, firstProfit[i]+secondProfit[i])
	}
	return max
}

//DP
//题意是最多2次交易的最大获利值，可以0，1，2次交易
//所以分成5个阶段，1，3，5是买入，卖出/买入，卖出。2，4是持有股票。每个阶段又分两种情况，如下
func BestTimeStock32(prices []int) int {
	n := len(prices)

	//i表示天数，j表示阶段。最多交易两次被表示为5个阶段
	//dp[i][j]代表前i天第j阶段获得的最大收益
	dp := [][]int{}
	for i := 0; i <= n; i++ {
		tmp := make([]int, 5+1)
		dp = append(dp, tmp)
	}

	for i := 1; i <= n; i++ {
		//phase1,3,5。现在手头没股票
		for j := 1; j <= 5; j += 2 {
			//前一天也没有股票
			dp[i][j] = dp[i-1][j]
			//前一天有股票，今天卖了。求前一天没有股票情况下的收益大还是前一天有股票情况下收益大
			//前一天有股票，证明不是第一阶段，所以j>1
			if i > 1 && j > 1 {
				dp[i][j] = infra.IntMax(dp[i][j], dp[i-1][j-1]+prices[i-1]-prices[i-2]) //prices是0base的
			}

		}
		//phase2,4。现在手头有股票
		for j := 2; j <= 5; j += 2 {
			//前一天无股票，今天买了股票
			dp[i][j] = dp[i-1][j-1]
			//前一天有股票，今天继续持有
			if i > 1 {
				dp[i][j] = infra.IntMax(dp[i][j], dp[i-1][j-1]+prices[i-1]-prices[i-2]) //prices是0base的
			}

		}
	}

	var res int
	for j := 1; j <= 5; j++ {
		res = infra.IntMax(res, infra.IntMax(0, dp[n][j]))
	}

	return res
}
