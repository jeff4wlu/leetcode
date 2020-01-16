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
