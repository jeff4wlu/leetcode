package leet7

import (
	"leetcode/base"
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
	if int(math.Abs(float64(res))) > base.INT32_MAX {
		return 0
	}

	return res
}
