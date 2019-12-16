package leet7

import (
	"leetcode/infra"
	"math"
)

func ReverseInt(x int) int {

	var res int

	for {
		if x == 0 {
			break
		}

		res = res*10 + x%10
		x /= 10
	}

	//int32溢出检查
	if int(math.Abs(float64(res))) > infra.INT32_MAX {
		return 0
	}

	return res
}
