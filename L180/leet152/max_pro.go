package leet152

import "leetcode/infra"

//暴力破解，最直接的思路
func MaxPro1(in []int) int {
	n := len(in)

	max := infra.INT32_MIN
	for i := 0; i < n; i++ {
		tmp := in[i]
		for j := i + 1; j < n; j++ {
			tmp *= in[j]
			max = infra.IntMax(max, tmp)
		}
	}

	return max
}

//DP，具有记忆功能
func MaxPro2(in []int) int {
	n := len(in)

	max := in[0]
	min := in[0]
	res := in[0]
	for i := 1; i < n; i++ {
		max = infra.IntMax(infra.IntMax(min*in[i], max*in[i]), in[i])
		min = infra.IntMin(infra.IntMin(min*in[i], max*in[i]), in[i])
		res = infra.IntMax(max, res)
	}

	return res
}
