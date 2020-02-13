package leet254

import "leetcode/infra"

func FactorCombination(num int) [][]int {

	if num < 4 || infra.IsPrime(num) {
		return [][]int{}
	}

	res := [][]int{}
	path := []int{}
	dfs(num, path, &res)
	return res
}

func dfs(num int, path []int, res *[][]int) {

	n := len(path)
	var start int
	if n > 0 {
		if num < path[n-1] {
			return
		}
		start = path[n-1]
	} else {
		start = 2
	}

	for i := start; i <= num; i++ {
		a := num / i
		b := num % i
		if b == 0 && a != 1 && a >= i {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, i)
			tmp = append(tmp, a)
			(*res) = append((*res), tmp)

			tmp = make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, i)
			dfs(a, tmp, res)

		}
	}

}
