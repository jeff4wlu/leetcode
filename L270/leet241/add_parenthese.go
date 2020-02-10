package leet241

import "strconv"

//加括号相当于去掉符号的优先级，变成排列组合问题
//学会dfs中返回多个结果，保存多个结果计算
func AddParenthese(exp string) []int {

	expS := []string{}
	for i := 0; i < len(exp); {
		if exp[i] == ' ' {
			i++
			continue
		} else if exp[i] == '*' || exp[i] == '+' || exp[i] == '-' {
			expS = append(expS, string(exp[i]))
			i++
			continue
		}
		var num string
		for ; i < len(exp) && exp[i] >= '0' && exp[i] <= '9'; i++ {
			num += string(exp[i])
		}
		expS = append(expS, num)
	}

	return dfs(expS, 0, len(expS)-1)

}

func dfs(expS []string, left, right int) []int {

	res := []int{}
	if left == right {
		tmp, _ := strconv.Atoi(string(expS[left]))
		return append(res, tmp)
	}

	for i := left + 1; i < right; i += 2 {
		larr := dfs(expS, left, i-1)
		rarr := dfs(expS, i+1, right)

		for _, m := range larr {
			for _, n := range rarr {
				if expS[i] == "+" {
					res = append(res, m+n)
				} else if expS[i] == "-" {
					res = append(res, m-n)
				} else {
					res = append(res, m*n)
				}

			}
		}
	}

	return res
}
