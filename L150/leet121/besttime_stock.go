package leet121

import "leetcode/infra"

//注意题目要求，不能在买入前卖出股票
func BestTimeStock(in []int) int {

	min := infra.INT32_MAX
	maxProfit := 0
	for i := 0; i < len(in); i++ {
		if in[i] < min {
			min = in[i]
		}
		if in[i]-min > maxProfit {
			maxProfit = in[i] - min
		}
	}

	return maxProfit

}
