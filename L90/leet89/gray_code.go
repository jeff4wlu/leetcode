package leet89

import "math"

import "leetcode/infra"

func GrayCode(n int) []int {

	var first string
	for i := 0; i < n; i++ {
		first += "0"
	}

	dict := map[string]int{}
	res := []string{}

	total := math.Exp2(float64(n))
	solve(first, n, int(total), dict, &res)

	resNums := []int{}
	for i := 0; i < len(res); i++ {
		num := infra.Str2DEC(res[i])
		resNums = append(resNums, num)
	}
	return resNums
}

func solve(in string, n, total int, dict map[string]int, res *[]string) {

	if len(dict) >= total {
		return
	}

	for i := 0; i < n; i++ {

		//改变string中某个元素的写法
		var s2 string
		tmp := []byte(in)
		if tmp[i] == '0' {
			tmp[i] = '1'
		} else {
			tmp[i] = '0'
		}
		s2 = string(tmp)
		if _, ok := dict[s2]; !ok {
			dict[s2] = 1
			*res = append(*res, s2)
			solve(s2, n, total, dict, res)
		}
	}

}
