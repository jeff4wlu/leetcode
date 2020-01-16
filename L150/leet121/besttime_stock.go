package leet121

import "leetcode/infra"

func BestTimeStock(in []int) int {

	min := infra.INT32_MAX
	max := infra.INT32_MIN
	for i := 0; i < len(in); i++ {
		if in[i] > max {
			max = in[i]
		}
		if in[i] < min {
			min = in[i]
		}
	}

	return max - min

}
