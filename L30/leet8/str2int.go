package leet8

import (
	"leetcode/infra"
)

func StrToInt(s string) int {

	if s == "" {
		return 0
	}

	var data = []byte(s)
	length := len(data)

	//去掉前置空格
	var start int
	for i := 0; i < length; i++ {
		if data[i] == ' ' {
			continue
		}
		start = i
		break
	}

	sign := 1
	if start < length && data[start] == '-' {
		sign = -1
		start++
	} else if start < length && data[start] == '+' {
		start++
	}

	var base int
	//判断是否数字，非数字跳出循环
	for i := start; i < length; i++ {
		if data[i] >= '0' && data[i] <= '9' {
			//int32溢出检查
			if base > infra.INT32_MAX {
				if sign > 0 {
					return infra.INT32_MAX
				}

				return infra.INT32_MIN
			}

			base = base*10 + int(data[i]) - int('0')
			continue

		}

		break

	}

	//单个符号无数字、前置有字符都返回0
	return base * sign

}
