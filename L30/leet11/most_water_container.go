package leet11

import "leetcode/infra"

func MostWaterContainer(in []int) int {

	var i, res int
	j := len(in) - 1
	for i < j {
		res = infra.IntMax(res, infra.IntMin(in[i], in[j])*(j-i))
		if in[i] <= in[j] {
			i++
		} else {
			j--
		}
	}

	return res
}
