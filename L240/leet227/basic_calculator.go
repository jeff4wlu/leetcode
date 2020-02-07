package leet227

import "leetcode/infra"

//继续利用stack来保存未用到的过程状态
//关键点是stack的数据存储结构是先符号，后数字，这个规则不能变。所以首push是符号+
func BasicCal(exp string) int {

	n := len(exp)
	sign := 1
	var res int

	stack := infra.CreateStack()
	stack.Push(1)
	for i := 0; i < n; i++ {
		tmp := exp[i]

		if tmp >= '0' && tmp <= '9' {
			num := int(tmp - '0')
			for ; i+1 < n && (exp[i+1] >= '0' && exp[i+1] <= '9'); i++ {
				num = num*10 + int((exp[i+1] - '0'))
			}
			if sign >= -1 {
				stack.Push(num)
				continue
			} else if sign == -2 {
				a := stack.Pop().(int)
				stack.Push(a * num)

			} else if sign == -3 {
				a := stack.Pop().(int)
				stack.Push(a / num)
			}
			sign = 1
			continue

		} else if tmp == '+' {
			stack.Push(1)
		} else if tmp == '-' {
			stack.Push(-1)
		} else if tmp == '*' {
			sign = -2
		} else if tmp == '/' {
			sign = -3
		}
	}

	res = 0
	for !stack.IsEmpty() {
		tmp := stack.Pop().(int)
		sign = stack.Pop().(int)
		res += tmp * sign
	}

	return res
}
