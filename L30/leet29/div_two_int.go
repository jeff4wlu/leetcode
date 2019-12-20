package leet29

import (
	"leetcode/infra"
	"math"
)

func DivTwoInt(dividend, divisor int) int {

	if dividend == infra.INT32_MIN && divisor == -1 {
		return infra.INT32_MAX
	}

	var ans int
	num1 := int(math.Abs(float64(dividend)))
	num2 := int(math.Abs(float64(divisor)))

	for num1 > num2 {
		var cnt int
		for num1 > num2<<cnt {
			cnt++
		}
		ans += (1 << (cnt - 1))
		num1 -= (num2 << (cnt - 1))
	}

	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 {
		return ans
	}

	return -ans

}
