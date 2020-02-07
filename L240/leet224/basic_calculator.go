package leet224

import "leetcode/infra"

//sign变量保存当前符号，res保存当前和值
//当遇到括号，要压栈保存当前状态，即sign和res变量值
func BasicCal(exp string) int {

	n := len(exp)

	var res int
	sign := 1
	stack := infra.CreateStack()

	for i := 0; i < n; i++ {
		tmp := exp[i]

		if tmp >= '0' && tmp <= '9' {
			num := int(tmp - '0')
			for ; i+1 < n && (exp[i+1] >= '0' && exp[i+1] <= '9'); i++ {
				num = num*10 + int((exp[i+1] - '0'))
			}
			res += num * sign
		} else if tmp == '+' {
			sign = 1
		} else if tmp == '-' {
			sign = -1
		} else if tmp == '(' {
			stack.Push(res)
			stack.Push(sign)
			res = 0
			sign = 1
		} else if tmp == ')' {
			res = res*stack.Pop().(int) + stack.Pop().(int)
		}
	}

	return res
}
